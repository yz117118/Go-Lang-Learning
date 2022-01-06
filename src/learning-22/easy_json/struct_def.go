package json

//内置json解析
//利用反射实现，通过FieldTag来标识对应的json值

type BasicInfo struct {
	Name string `json: "name"`
	Age  int    `json: "age"`
}

type JobInfo struct {
	Skills []string `json: "skills"`
}

type Employee struct {
	BasicInfo BasicInfo `json:"basic_info"`
	JobInfo   JobInfo   `json:"job_info"`
}
