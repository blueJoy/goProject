package main

import (
	"net/http"
	"io"
	"os"
	"io/ioutil"
	"log"
	"html/template"
	"runtime/debug"
	"path"
)

const (
	WEB_ROOT = "src\\photoweb"
	UPLOAD_DIR = WEB_ROOT+"\\uploads"
	TEMPLATE_DIR=WEB_ROOT+"\\template"

	ListDir = 0x0001
)

//缓存模板
var templates map[string]*template.Template

var tmplNames = []string{"upload","list","err"}


func uploadHandler(w http.ResponseWriter,r *http.Request){
	if r.Method == "GET"{
		//修改成HTML模板
		/*io.WriteString(w,"<body><form method=\"POST\" action=\"/upload\" "+
			" enctype=\"multipart/form-data\">"+
			"Choose an image to upload: <input name=\"image\" type=\"file\" />"+
			"<input type=\"submit\" value=\"Upload\" />"+
			"</form></body>")*/

		/*	t,err := template.ParseFiles(TEMPLATE_DIR+"\\"+"upload.html")
			if err != nil{
				http.Error(w,err.Error(),http.StatusInternalServerError)
				return
			}
			t.Execute(w,nil)*/

			renderHtml(w,"upload",nil)
		return
	}

	if r.Method == "POST"{
		f,h,err := r.FormFile("image")
	/*	if err != nil{
			http.Error(w,err.Error(),http.StatusInternalServerError)
			return
		}*/
		check(err)
		filename := h.Filename
		defer f.Close()

		//makedir
		/*_,err1 := os.Stat(UPLOAD_DIR)
		if err1 != nil && os.IsNotExist(err1){
			os.Mkdir(UPLOAD_DIR,os.ModePerm)
		}*/
		if isExist(UPLOAD_DIR){
			os.Mkdir(UPLOAD_DIR,os.ModePerm)
		}

		//createFile
		t,err := os.Create(UPLOAD_DIR+"/"+filename)
	/*	if err != nil{
			http.Error(w,err.Error(),http.StatusInternalServerError)
			return
		}*/
		check(err)
		defer t.Close()

		if _,err := io.Copy(t,f); err != nil{
			http.Error(w,err.Error(),http.StatusInternalServerError)
			return
		}

		http.Redirect(w,r,"/view?id="+filename,http.StatusFound)
	}
}

func viewHandler(w http.ResponseWriter,r *http.Request){
	imageId := r.FormValue("id")
	imagePath := UPLOAD_DIR + "/"+imageId
	if !isExist(imagePath){
		http.NotFound(w,r)
		return
	}

	w.Header().Set("Content-Type","image")
	http.ServeFile(w,r,imagePath)
}

func isExist(path string) bool{
	_,err := os.Stat(path)
	if err ==nil{
		return true
	}
	return os.IsExist(err)
}

func listHandler(w http.ResponseWriter,r *http.Request)  {
	fileInfoArr,err := ioutil.ReadDir(UPLOAD_DIR)
	/*if err != nil{
		http.Error(w,err.Error(),http.StatusInternalServerError)
		return
	}*/
	check(err)

	//修改成用模板
	/*var listHtml string

	for _,fileInfo := range fileInfoArr {
		imgid := fileInfo.Name()
		listHtml += "<li><a href=\"/view?id="+imgid+"\">"+imgid+"</a></li>"
	}
	io.WriteString(w,"<body><ol>"+listHtml+"</ol></body>")*/

	locals := make(map[string]interface{})
	images := []string{}
	for _,fileInfo := range  fileInfoArr{
		images = append(images,fileInfo.Name())
	}
	locals["images"] = images
	/*t,err := template.ParseFiles(TEMPLATE_DIR+"\\"+"list.html")
	if err != nil{
		http.Error(w,err.Error(),http.StatusInternalServerError)
		return
	}
	t.Execute(w,locals)*/

	renderHtml(w,"list",locals)
}

//模板处理方法
func renderHtml(w http.ResponseWriter,tmpl string ,locals map[string]interface{}){

	/*t,err := template.ParseFiles(TEMPLATE_DIR+"\\"+tmpl+".html")
	if err != nil{
		http.Error(w,err.Error(),http.StatusInternalServerError)
		return
	}
	t.Execute(w,locals)*/

	//使用缓存
	if err := templates[tmpl].Execute(w,locals);err != nil{
		http.Error(w,err.Error(),http.StatusInternalServerError)
		return
	}

}

//初始化缓存HTML模板
func initTemplate() {

	templates = make(map[string]*template.Template)

	//tmplName 模板名字用数组有限制，可以改成配置文件，配置中心，或者读模板目录下的文件形式等
	for _,tmpl := range tmplNames{
		t := template.Must(template.ParseFiles(TEMPLATE_DIR+"\\"+tmpl+".html"))
		templates[tmpl] = t
	}
}

func initTemplateByFile(){
	fileInfoArr ,err := ioutil.ReadDir(TEMPLATE_DIR)
	if err != nil{
		panic(err)
	}
	var templateName,templatePath string
	for _,fileInfo := range fileInfoArr{
		templateName = fileInfo.Name()
		if ext := path.Ext(templateName); ext != ".html"{
			continue
		}
		templatePath = TEMPLATE_DIR + "\\" + templateName
		//截取名字
		tmpl := templateName[:len(".html")]
		log.Println("Loading template:",templatePath)
		t := template.Must(template.ParseFiles(templatePath))
		templates[tmpl] = t
	}

}

//该中方式如果出错会导致崩溃
func check(err error){
	if err != nil{
		panic(err)
	}
}

//错误集中处理
func safeHandler(fn http.HandlerFunc) http.HandlerFunc{

	//巧妙地使用了 defer 关键字搭配 recover() 方法终结 panic 的肆行

	return func(w http.ResponseWriter,r *http.Request){
		defer func(){
			if e,ok := recover().(error);ok{

				//http.Error(w,e.Error(),http.StatusInternalServerError)

				//或者数组自定义的50x错误页面
				locals := make(map[string]interface{})
				locals["err"] = e.Error()
				renderHtml(w,"err",locals)
				log.Println("WARN: panic in %v - %v", fn, e)
				log.Println(string(debug.Stack()))
			}
		}()
		fn(w,r)
	}
}

//静态资源暴露
func staticDirHandler(mux *http.ServeMux,prefix string,staticDir string,flag int){
	mux.HandleFunc(prefix, func(w http.ResponseWriter, r *http.Request) {
		file :=  staticDir + r.URL.Path[len(prefix)-1:]
		if (flag & ListDir == 0){
			if exists := isExists(file); !exists{
				http.NotFound(w,r)
				return
			}
		}
		http.ServeFile(w,r,file)
	})
}
//判断文件是否存在
func isExists(path string) bool{
	_,err := os.Stat(path)
	if err == nil{
		return true
	}
	return os.IsNotExist(err)
}


func main (){

	//初始化HTML模板
	initTemplate()



/*	http.HandleFunc("/upload",uploadHandler)
	http.HandleFunc("/view",viewHandler)
	http.HandleFunc("/",listHandler)	*/

	//使用默认的DefaultServerMux
/*	http.HandleFunc("/upload",safeHandler(uploadHandler))
	http.HandleFunc("/view",safeHandler(viewHandler))
	http.HandleFunc("/",safeHandler(listHandler))

	err := http.ListenAndServe(":8080",nil)*/

	//暴露静态资源
	mux := http.NewServeMux()
	staticDirHandler(mux,"/assets/",WEB_ROOT+"/public",0)

	mux.HandleFunc("/upload",safeHandler(uploadHandler))
	mux.HandleFunc("/view",safeHandler(viewHandler))
	mux.HandleFunc("/",safeHandler(listHandler))

	err := http.ListenAndServe(":8080",mux)

	if err !=nil{
		log.Fatal("ListenAndServe: ",err.Error())
	}

}
