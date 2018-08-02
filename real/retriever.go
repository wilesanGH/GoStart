package real

type Retriever interface {

	Get(url string) string
}


func Download(r Retriever) string{
	return r.Get("http://cec.usst.edu.cn")
}

