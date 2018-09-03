package engine

type ConcurrentEngine struct {
	Scheduler Scheduler
	WorkerCount int
	ItemChan chan interface{}
}

type Scheduler interface {
	ReadyNotifer
	Submit(Request)
	WorkerChan() chan Request
	Run()
}

type ReadyNotifer interface {
	WorkerReady(chan Request)
}

func (e *ConcurrentEngine) Run(seeds ...Request)  {

	out := make(chan ParseResult)
	e.Scheduler.Run()
	for i := 0 ;i< e.WorkerCount;i++{
		createWorker(e.Scheduler.WorkerChan(),out,e.Scheduler)
	}

	for _, r:=range seeds{
		e.Scheduler.Submit(r)
	}

	//itemCount := 0

	for{
		result := <- out
		for _,item := range result.Items{
			//itemCount++
			//log.Printf("GOT item: %v,itemCount %d",item,itemCount)
			go func() {e.ItemChan <- item}()
		}

		for _, request := range result.Requests{
			e.Scheduler.Submit(request)
		}
	}
}


func createWorker(in chan Request,out chan ParseResult,ready ReadyNotifer){
	go func() {
		for {
			// tell scheduler i'm ready
			ready.WorkerReady(in)
			request := <-in
			result, err := Worker(request)
			if err != nil {
				continue
			}
			out <- result
		}
	}()
}