package sparks

import (
   "fmt"
   "testing"
   "time"
)

func InitData(stopCh <-chan struct{}, funcOptions ...initFunc) {
   for _, opt := range funcOptions {
      go RunUntil(opt.InitFunc, stopCh, opt.Duration)
   }
}

func RunUntil(f func(), stopCh <-chan struct{}, duration time.Duration) {
   ticker := time.NewTicker(duration)

   for {
      select {
      case <-stopCh:
         fmt.Println("th first stopCh")
         return
      default:
      }

      f()

      select {
      case <-ticker.C:
         fmt.Println("time.ticker")
      case <-stopCh:
         fmt.Println("the second stopCh")
         return
      }
   }
}

func Shutdown(stopCh chan struct{}) {
   close(stopCh)
}

func InitData1() {
   for i :=1; i < 10; i++ {
      time.Sleep(1 * time.Second)
      fmt.Printf("sleesp %d\n", i)
   }

   fmt.Println("initData1")
}

func InitData2() {
   fmt.Println("init data 2")
}

func InitData4() {
   fmt.Println("init data 4")
}

type initFunc struct {
   InitFunc func()
   Duration time.Duration
}

func TestRountinControl(t *testing.T) {
   fmt.Println("asf")

   stopCh := make(chan struct{})

   initFuncOpts := []initFunc{
      {InitData1, time.Duration(1) * time.Second},
      {InitData4, time.Duration(3) * time.Second},
      {InitData2, time.Duration(5) * time.Second},
   }

   InitData(stopCh, initFuncOpts...)

   for i := 0; i < 10; i++ {
      time.Sleep(1 * time.Second)
      fmt.Println("sleep")
   }

   Shutdown(stopCh)
}

func TestX(t *testing.T)  {
   stopCh := make(chan struct{}, 1)

   for i:= 0; i < 3 ; i++ {
      go func() {
         select {
         case v:=<- stopCh:
            fmt.Printf("%d get %v\n", i, v)
            return
         default:
            time.Sleep(1)
         }
      }()
   }
   time.Sleep(10)
   stopCh <- struct{}{}
}

