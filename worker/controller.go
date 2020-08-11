package worker



type controller struct {
	workers 		map[int]*worker
	channels 		map[int](chan Response)
	

}

func InitController( num int) (c *controller){
	return c 
}