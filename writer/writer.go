package writer
import(
	"fmt"
	"time"
)

func Writer(jobs <-chan int, out chan<- string){
	fmt.Println("Starting writer with ", jobs)
	val := <-jobs
	fmt.Println("received ", val)
	time.Sleep(time.Second * 5)
	fmt.Println("finished processing ", val)
	out <- "finished processing "
}
