package main

import (
	
   "fmt"
   helper "./helpers"
   "github.com/gin-contrib/cors"                        // Why do we need this package?
   "github.com/gin-gonic/gin"
   "github.com/jinzhu/gorm"
   _ "github.com/jinzhu/gorm/dialects/sqlite"           // If you want to use mysql or any other db, replace this line
)

var db *gorm.DB                                         // declaring the db globally
var err error

type Person struct {
   ID uint `json:"id"`
   FirstName string `json:"firstname"`
   LastName string `json:"lastname"`
   //City string `json:"city"`
   Username string  `json:"username"`
   Email string  `json:"email"` 
   Password string `json:"password"`
   Admin uint `json:"admin"`
   Totalpoints uint `json:"totalpoints"`
}

type Quiz struct {
    ID uint `json:"id"`
    Genrename string `json:"genrename"`
    Quiznumber uint `json:"quiznumber"`
}

type Question struct {
    ID uint `json:"id"`
    Genrename string `json:"genrename"`
    Quiznumber uint `json:"quiznumber"`
    Type uint `json:"type"`
    Question string `json:"question"`
    Option1 string `json:"option1"`
    Option2 string `json:"option2"`
    Option3 string `json:"option3"`
    Option4 string `json:"option4"`
    Answer1 bool `json:"answer1"`
    Answer2 bool `json:"answer2"`
    Answer3 bool `json:"answer3"`
    Answer4 bool `json:"answer4"`
}

type Answers struct {
    Answer1 bool `json:"answer1"`
    Answer2 bool `json:"answer2"`
    Answer3 bool `json:"answer3"`
    Answer4 bool `json:"answer4"`
}

type Quizpoints struct {
    ID uint `json:"id"`  
    Username string  `json:"username"`
    Genrename string `json:"genrename"`
    Quiznumber uint `json:"quiznumber"`
    Points uint `gorm:"default:0" , json:"points"`
}

type Genrepoints struct {
    ID uint `json:"id"`
    Username string  `json:"username"`
    Genrename string `json:"genrename"`
    Gpoints uint `gorm:"default:0" , json:"gpoints"`
}




func main() {
   db, err = gorm.Open("sqlite3", "./gorm.db")
   if err != nil {
      fmt.Println(err)
   }
   defer db.Close()                                          //to close it 

   db.AutoMigrate(&Person{} , &Quiz{} , &Question{} ,&Quizpoints{} , &Genrepoints{})

//    var person1 = Person{FirstName: "Zega", LastName: "Tron" , Username: "zt" , Email:"zt@",Password:"pass",Admin:0,Totalpoints:0}
//    db.Create(&person1)
//    var person2 = Person{FirstName: "admin", LastName: "admin" , Username: "admin" , Email:"admin@",Password:"pass",Admin:1,Totalpoints:0}
//    db.Create(&person2)
//    var person3 = Person{FirstName: "aj", LastName: "aj" , Username: "aj" , Email:"aj@",Password:"pass",Admin:0,Totalpoints:0}
//    db.Create(&person3)
//    var person4 = Person{FirstName: "dj", LastName: "dj" , Username: "dj" , Email:"dj@",Password:"pass",Admin:0,Totalpoints:0}
//    db.Create(&person4)

//    var quiz1 = Quiz{Genrename:"HP" , Quiznumber:1}
//    db.Create(&quiz1)
//    var quiz2 = Quiz{Genrename:"HP" , Quiznumber:2}
//    db.Create(&quiz2)
//    var quiz3 = Quiz{Genrename:"HP" , Quiznumber:3}
//    db.Create(&quiz3)
//    var quiz4 = Quiz{Genrename:"SH" , Quiznumber:1}
//    db.Create(&quiz4)
//    var quiz5 = Quiz{Genrename:"SH" , Quiznumber:2}
//    db.Create(&quiz5)
//    var quiz6 = Quiz{Genrename:"Physics" , Quiznumber:1}
//    db.Create(&quiz6)
//    var quiz7 = Quiz{Genrename:"Physics" , Quiznumber:2}
//    db.Create(&quiz7)
//    var quiz8 = Quiz{Genrename:"SH" , Quiznumber:3}
//    db.Create(&quiz8)

//    var question1 = Question{Genrename: "Physics", Quiznumber : 1 , Type : 0 ,Question : "What is the value of 'g'?(Values are in m/((sec)^2))",Option1 : "10" , Option2 : "20" ,Option3 : "9.8" , Option4 : "9" , Answer1 : false ,Answer2 : false , Answer3 : true , Answer4 : false}
//    db.Create(&question1)
//    var question2 = Question{Genrename: "Physics", Quiznumber : 1 , Type : 1 ,Question : "Who among the following Discovered Radioactivity ?",Option1 : "Marie Curie" , Option2 : "Pierre Curie" ,Option3 : "Bequerl" , Option4 : "Bohr" , Answer1 : true ,Answer2 : true , Answer3 : true , Answer4 : false}
//    db.Create(&question2)
//    var question3 = Question{Genrename: "Physics", Quiznumber : 1 , Type : 0 ,Question : "What is the Second Law Of Motion",Option1 : "V=u+at" , Option2 : "F=m*a" ,Option3 : "has not been discovered" , Option4 : "F=(G*m1*m2)/(r*r)" , Answer1 : false ,Answer2 : true , Answer3 : false , Answer4 : false}
//    db.Create(&question3)
//    var question4 = Question{Genrename: "Physics", Quiznumber : 1 , Type : 0 ,Question : "The Shape of our Milky Way Galaxy is",Option1 : "circular" , Option2 : "elliptical" ,Option3 : "spiral" , Option4 : "none" , Answer1 : false ,Answer2 : false , Answer3 : true , Answer4 : false }
//    db.Create(&question4)
//    var question5 = Question{Genrename: "Physics", Quiznumber : 1 , Type : 0 ,Question : "Who are scientists?",Option1 : "Einstien" , Option2 : "Newton" ,Option3 : "Classmate" , Option4 : "NOne" , Answer1 : true ,Answer2 : true , Answer3 : false , Answer4 : false}
//    db.Create(&question5)

//    var question6 = Question{Genrename: "Physics", Quiznumber : 2 , Type : 1 ,Question : "Who are scientists ?",Option1 : "Einstien" , Option2 : "Newton" ,Option3 : "Bohr" , Option4 : "NOne" , Answer1 : true ,Answer2 : true , Answer3 : true , Answer4 : false}
//    db.Create(&question6)
//    var question7 = Question{Genrename: "Physics", Quiznumber : 2 , Type : 0 ,Question : "Who are scientists2?",Option1 : "Planck" , Option2 : "Bunny" ,Option3 : "Classmate" , Option4 : "one" , Answer1 : true ,Answer2 : false , Answer3 : false , Answer4 : false}
//    db.Create(&question7)
//    var question8 = Question{Genrename: "Physics", Quiznumber : 2 , Type : 0 ,Question : "Who are scientists3?",Option1 : "Hisenberg" , Option2 : "Newton Jr." ,Option3 : "Classmate" , Option4 : "None" , Answer1 : true ,Answer2 :false , Answer3 : false , Answer4 : false}
//    db.Create(&question8)
//    var question9 = Question{Genrename: "Physics", Quiznumber : 2 , Type : 1 ,Question : "Who are scientists4?",Option1 : "Max Born" , Option2 : "Newton 5" ,Option3 : "Classmates" , Option4 : "Young" , Answer1 : true ,Answer2 : false , Answer3 : false , Answer4 : true}
//    db.Create(&question9)

//    var question10 = Question{Genrename: "HP", Quiznumber : 1 , Type : 0 ,Question : "How many books does HP series have?",Option1 : "1" , Option2 : "10" ,Option3 : "7" , Option4 : "none" , Answer1 : false ,Answer2 : false , Answer3 : true , Answer4 : false }
//    db.Create(&question10)
//    var question11 = Question{Genrename: "HP", Quiznumber : 1 , Type : 0 ,Question : "How many movies HP series have?",Option1 : "8" , Option2 : "10" ,Option3 : "7" , Option4 : "none" , Answer1 : true ,Answer2 : false , Answer3 : false , Answer4 : false }
//    db.Create(&question11)
//    var question12 = Question{Genrename: "HP", Quiznumber : 1 , Type : 1 ,Question : "Harry's Best Friends?",Option1 : "Malfoy" , Option2 : "Ron" ,Option3 : "Hermione" , Option4 : "Sirius" , Answer1 : false ,Answer2 : true, Answer3 : true , Answer4 : true }
//    db.Create(&question12)
//    var question13 = Question{Genrename: "HP", Quiznumber : 1 , Type : 1 ,Question : "Ron's Best Friends?",Option1 : "Malfoy" , Option2 : "Harry" ,Option3 : "Hermione" , Option4 : "Sirius" , Answer1 : false ,Answer2 : true, Answer3 : true , Answer4 : false }
//    db.Create(&question13)
//    var question14 = Question{Genrename: "HP", Quiznumber : 1 , Type : 0 ,Question : "Voldermort's Best Friends?",Option1 : "Malfoy" , Option2 : "Ron" ,Option3 : "Hermione" , Option4 : "Sirius" , Answer1 : true ,Answer2 : false, Answer3 : false , Answer4 : false }
//    db.Create(&question14)

//    var question15 = Question{Genrename: "SH", Quiznumber : 1 , Type : 0 ,Question : "Sherlock's Best Friends?",Option1 : "Malfoy" , Option2 : "Ron" ,Option3 : "Watson" , Option4 : "Sirius" , Answer1 : false ,Answer2 : false, Answer3 : true , Answer4 : false }
//    db.Create(&question15)
//    var question16 = Question{Genrename: "SH", Quiznumber : 1 , Type : 1 ,Question : "Sherlock's Enemy?",Option1 : "Sherlock's Sister" , Option2 : "Moriarity" ,Option3 : "Watson" , Option4 : "Voldermort" , Answer1 : true ,Answer2 : true, Answer3 : false , Answer4 : false }
//    db.Create(&question16)
//    var question17 = Question{Genrename: "SH", Quiznumber : 1 , Type : 0 ,Question : "Sherlock's Son?",Option1 : "None" , Option2 : "Moriarity" ,Option3 : "Watson" , Option4 : "Voldermort" , Answer1 : true ,Answer2 : false, Answer3 : false , Answer4 : false }
//    db.Create(&question17)
//    var question18 = Question{Genrename: "SH", Quiznumber : 1 , Type : 0 ,Question : "Sherlock's Landlord?",Option1 : "Mrs. Hudson" , Option2 : "Moriarity" ,Option3 : "Watson" , Option4 : "Voldermort" , Answer1 : true ,Answer2 : false, Answer3 : false , Answer4 : false }
//    db.Create(&question18)
//    var question19 = Question{Genrename: "SH", Quiznumber : 1 , Type : 0 ,Question : "Sherlock's Address?",Option1 : "221 B Down Town" , Option2 : "221 B China Town" ,Option3 : "221 B Baker Town" , Option4 : "221 B Baker Street" , Answer1 : false ,Answer2 : false, Answer3 : false , Answer4 : true }
//    db.Create(&question19)


//    var question20 = Question{Genrename: "SH", Quiznumber : 2 , Type : 0 ,Question : "Sherlock's Address?",Option1 : "221 B Down Town" , Option2 : "221 B China Town" ,Option3 : "221 B Baker Town" , Option4 : "221 B Baker Street" , Answer1 : false ,Answer2 : false, Answer3 : false , Answer4 : true }
//    db.Create(&question20)
//    var question21 = Question{Genrename: "SH", Quiznumber : 2 , Type : 0 ,Question : "Sherlock's Address?",Option1 : "221 B Down Town" , Option2 : "221 B China Town" ,Option3 : "221 B Baker Town" , Option4 : "221 B Baker Street" , Answer1 : false ,Answer2 : false, Answer3 : false , Answer4 : true }
//    db.Create(&question21)
//    var question22 = Question{Genrename: "SH", Quiznumber : 2 , Type : 0 ,Question : "Sherlock's Address?",Option1 : "221 B Down Town" , Option2 : "221 B China Town" ,Option3 : "221 B Baker Town" , Option4 : "221 B Baker Street" , Answer1 : false ,Answer2 : false, Answer3 : false , Answer4 : true }
//    db.Create(&question22)

//    var question23 = Question{Genrename: "SH", Quiznumber : 3 , Type : 0 ,Question : "Sherlock's Best Friends?",Option1 : "Malfoy" , Option2 : "Ron" ,Option3 : "Watson" , Option4 : "Sirius" , Answer1 : false ,Answer2 : false, Answer3 : true , Answer4 : false }
//    db.Create(&question23)
//    var question24 = Question{Genrename: "SH", Quiznumber : 3 , Type : 0 ,Question : "Sherlock's Best Friends?",Option1 : "Malfoy" , Option2 : "Ron" ,Option3 : "Watson" , Option4 : "Sirius" , Answer1 : false ,Answer2 : false, Answer3 : true , Answer4 : false }
//    db.Create(&question24)
//    var question25 = Question{Genrename: "SH", Quiznumber : 3 , Type : 0 ,Question : "Sherlock's Best Friends?",Option1 : "Malfoy" , Option2 : "Ron" ,Option3 : "Watson" , Option4 : "Sirius" , Answer1 : false ,Answer2 : false, Answer3 : true , Answer4 : false }
//    db.Create(&question25)

//    var question26 = Question{Genrename: "HP", Quiznumber : 2 , Type : 1 ,Question : "Harry's Best Friends2?",Option1 : "Malfoy" , Option2 : "Ron" ,Option3 : "Hermione" , Option4 : "Sirius" , Answer1 : false ,Answer2 : true, Answer3 : true , Answer4 : true }
//    db.Create(&question26)

//    var question27 = Question{Genrename: "HP", Quiznumber : 3 , Type : 1 ,Question : "Harry's Best Friends3?",Option1 : "Malfoy" , Option2 : "Ron" ,Option3 : "Hermione" , Option4 : "Sirius" , Answer1 : false ,Answer2 : true, Answer3 : true , Answer4 : true }
//    db.Create(&question27)
//    var question28 = Question{Genrename: "HP", Quiznumber : 3 , Type : 1 ,Question : "Harry's Best Friends3?",Option1 : "Malfoy" , Option2 : "Ron" ,Option3 : "Hermione" , Option4 : "Sirius" , Answer1 : false ,Answer2 : true, Answer3 : true , Answer4 : true }
//    db.Create(&question28)
   

   r := gin.Default()
   r.POST("/signup/" , SignUp)
   r.POST("/login/" , LogIn)
   r.POST("/createquiz/" , CreateQuiz)
   r.POST("/createquestion/",CreateQuestion)
   r.GET("/viewquizzes/",ViewQuizzes)
   r.GET("/viewquiz/:id",ViewQuiz)
   r.GET("/pastquizzes/:username" , PastQuizzes)
   r.POST("/evaluatequestion/:username/:id",EvaluateQuestion)
   r.GET("/leaderboard/",LeaderBoard)
   r.GET("/genreleaderboard/:genrename" , GenreLeaderBoard)
   r.DELETE("/deletequiz/:id" ,DeleteQuiz)
   r.DELETE("/deletequestion/:id",DeleteQuestion)
   r.GET("/adminmode/", AdminMode)
   r.GET("/people/", GetPeople)                             // Creating routes for each functionality
   r.GET("/people/:id", GetPerson)
   //r.POST("/people", CreatePerson)
   r.PUT("/people/:id", UpdatePerson)
   r.DELETE("/people/:id", DeletePerson)
   r.Use((cors.Default()))
   r.Run(":8080")                                           // Run on port 8080

}

// --------------------SIGNUP-----------------------------
func SignUp(c *gin.Context) {
    var person Person
    c.BindJSON(&person)

    var person2 Person
    var person3 Person

    //CHECK FOR EMPTY STRINGS
    firstname_chk := helper.IsEmpty(person.FirstName)
    lastname_chk := helper.IsEmpty(person.LastName)
    username_chk := helper.IsEmpty(person.Username)
    email_chk := helper.IsEmpty(person.Email)
    pass_chk := helper.IsEmpty(person.Password)

    if firstname_chk || lastname_chk || username_chk || email_chk || pass_chk {
        c.Header("access-control-allow-origin", "*")
        c.JSON(404, "Fill all fields!!!\n")
        fmt.Println("ErrorCode is -10 : There is empty data.")   //here you will have to return to front end if there is empty string
        return
    }

    buff_uname := person.Username
    buff_email := person.Email

    if s1 := db.Where("username = ?", buff_uname).First(&person2).Error; s1 == nil{
        c.Header("access-control-allow-origin", "*")
        c.JSON(404, "username exists!!\n")
        c.AbortWithStatus(404)
        fmt.Println("username exists!!")
        fmt.Println(s1)
        return
    }                                                                                           //what if there are multiple database, maybe it maps from where the person comes
    if s2 := db.Where("email = ?", buff_email).First(&person3).Error; s2 == nil{
        c.Header("access-control-allow-origin", "*")
        c.JSON(404, "email exists!!")
        c.AbortWithStatus(404)
        fmt.Println("email exists!!")
        fmt.Println(s2)
        return
    }                                                                                           //what if there are multiple database, maybe it maps from where the person comes
    
    db.Create(&person)
    
    if buff_uname == "admin" {
        person.Admin = 1
    }

    db.Save(&person)

    c.Header("access-control-allow-origin", "*")
    c.JSON(200,person)
    fmt.Println("Registration successful.")
}

// -----------------------------------LOGIN--------------------------------------
func LogIn(c *gin.Context){
    var person Person
    c.BindJSON(&person)
    var person2 Person

    username_chk := helper.IsEmpty(person.Username)
    pass_chk := helper.IsEmpty(person.Password)

    if username_chk || pass_chk {
        c.Header("access-control-allow-origin", "*")
        c.AbortWithStatus(404)
        c.JSON(404, "Fill all fields!!!\n")
        fmt.Println("ErrorCode is -10 : There is empty data.")   //here you will have to return to front end if there is empty string
        return
    }

    if s1 := db.Where("username = ?", person.Username).First(&person2).Error; s1 == nil{
        if person.Password == person2.Password{
            c.Header("access-control-allow-origin", "*") // Why am I doing this? Find out. Try running with this line commented
            c.JSON(200,person2)
            return        
        } else{
            c.Header("access-control-allow-origin", "*")
            c.AbortWithStatus(404)
            c.JSON(404 , "PASSWORD IS WRONG!")
            fmt.Println("WRONG PASS")
        }    
    } else {
        c.Header("access-control-allow-origin", "*")
        c.AbortWithStatus(404)
        c.JSON(404, "Username does not exist")
        fmt.Println(s1)
    }
}

func AdminMode(c *gin.Context){                       //from the frontend the person should be sent in .
    var person Person
    c.BindJSON(&person)


    if person.Admin==1 {
        c.JSON(200, "Welcome to admin mode\n")
    } else {
        c.JSON(404, "Sorry , you are not a admin!\n")
    }
}


func CreateQuiz(c *gin.Context){
    var quiz Quiz
    var quiz2 Quiz
    c.BindJSON(&quiz)

    genre_chk := helper.IsEmpty(quiz.Genrename)
    
    if genre_chk {
        c.Header("access-control-allow-origin", "*")
        c.JSON(404,"Genre Name Empty")
        fmt.Println("Genre Name Empty")
        return 
    }
    
    

    if err := db.Where("genrename = ?",quiz.Genrename).Where("quiznumber= ?",quiz.Quiznumber).First(&quiz2).Error; err == nil {
        c.Header("access-control-allow-origin", "*")
        c.JSON(404,"This quiz number exists for this genre")
        fmt.Println("This quiz number exists for this genre")
        return
    }

    db.Create(&quiz)
    c.Header("access-control-allow-origin", "*")
    c.JSON(200,"Quiz Created")

}

func CreateQuestion(c *gin.Context){
    var question Question
    c.BindJSON(&question)

    ques_chk := helper.IsEmpty(question.Question)
    opt1_chk := helper.IsEmpty(question.Option1)
    opt2_chk := helper.IsEmpty(question.Option2)
    opt3_chk := helper.IsEmpty(question.Option3)
    opt4_chk := helper.IsEmpty(question.Option4)


    if ques_chk || opt1_chk || opt2_chk || opt3_chk || opt4_chk {
        c.Header("access-control-allow-origin", "*")
        c.JSON(404,"Empty Field Encountered")
        fmt.Println("Field  empty")
        return
    }

    db.Create(&question)
    c.Header("access-control-allow-origin","*")
    c.JSON(200,"Question Created Successfully")
}

func ViewQuizzes(c *gin.Context) {
    var quiz []Quiz

    if err := db.Find(&quiz).Error ; err != nil {
        c.Header("access-control-allow-origin", "*")
        c.JSON(404,"NO DATA")
        fmt.Println(err)
    } else{
        c.Header("access-control-allow-origin", "*")
        c.JSON(200,quiz)    
    }

}

func ViewQuiz(c *gin.Context){
    id := c.Params.ByName("id")
    var genre string
    var qnum uint
    var quiz Quiz
    if err := db.Where("id = ?", id).First(&quiz).Error; err != nil {
        c.Header("access-control-allow-origin", "*")
        c.AbortWithError(404,err)
        fmt.Println(err)
     } else {
        //c.Header("access-control-allow-origin", "*") // Why am I doing this? Find out. Try running with this line commented
        //c.JSON(200, person)
        genre = quiz.Genrename
        qnum = quiz.Quiznumber    
        fmt.Println(genre)
        fmt.Println(qnum)
     }
    var questions []Question
    //.Where("quiznumber = ?", qnum)
    if err := db.Where("genrename = ?" , genre).Where("quiznumber = ?", qnum).Find(&questions).Error ; err != nil {
        c.Header("access-control-allow-origin", "*")
        c.JSON(404,"NO DATA")
        fmt.Println(err)
    } else{
        c.Header("access-control-allow-origin", "*")
        c.JSON(200,questions)    
    }
}

func EvaluateQuestion(c *gin.Context){
    user := c.Params.ByName("username")
    id := c.Params.ByName("id")
    var answers Answers
    c.BindJSON(&answers)
    var question Question

    var personQp Quizpoints
    var personGp Genrepoints
    var person Person

    if err := db.Where("id = ?",id).First(&question).Error ; err != nil {
        c.Header("access-control-allow-origin", "*")
        c.AbortWithError(404,err)
        return
    }

    chk_1 := (question.Answer1 == answers.Answer1)
    // fmt.Println("chek1" ,chk_1)
    chk_2 := (question.Answer2 == answers.Answer2)
    // fmt.Println(chk_2)
    chk_3 := (question.Answer3 == answers.Answer3)
    // fmt.Println(chk_3)
    chk_4 := (question.Answer4 == answers.Answer4)
    // fmt.Println("chk_4=",chk_4)


    if chk_1 && chk_2 && chk_3 && chk_4 { 

        db.Where("username = ?",user).First(&person)
        fmt.Println(person)
        person.Totalpoints += 5
        db.Save(&person)
        fmt.Println(person)

        if err1 := db.Where("username = ?",user).Where("genrename = ?" , question.Genrename).First(&personGp).Error ; err1 == nil {      /// if username ki row exist karti hai to err = nil and we just need to update otherwise create
            personGp.Gpoints += 5
            //fmt.Println(personGp)
            db.Save(&personGp)
            fmt.Println("Updated the value In Genre Table , Correct Answer")
        } else {
            // db.Create(&personQp)
            personGp.Username = user
            personGp.Genrename = question.Genrename
            personGp.Gpoints = personGp.Gpoints + 5
            db.Create(&personGp)
            fmt.Println("Created A New user in Genre Points Table  and updated score , Correct Answer")
        }

        if err := db.Where("username = ?",user).Where("genrename = ?" , question.Genrename).Where("quiznumber = ?",question.Quiznumber).First(&personQp).Error ; err == nil {      /// if username ki row exist karti hai to err = nil and we just need to update otherwise create
            personQp.Points += 5
            fmt.Println(personQp)
            db.Save(&personQp)
            fmt.Println("Updated the value , correct answer")
        } else {
            // db.Create(&personQp)
            personQp.Username = user
            personQp.Genrename = question.Genrename
            personQp.Quiznumber = question.Quiznumber
            personQp.Points = personQp.Points + 5
            db.Create(&personQp)
            fmt.Println("Created A New user and updated score in QuizPoints Table , Correct Answer")
        }
    } else {

        if err2 := db.Where("username = ?",user).Where("genrename = ?" , question.Genrename).First(&personGp).Error ; err2 == nil {      /// if username ki row exist karti hai to err = nil and we just need to update otherwise create
            fmt.Println("Wrong Answer , but record exists in Genre Table")
        } else {
            personGp.Username = user
            personGp.Genrename = question.Genrename
            db.Create(&personGp)
            fmt.Println("Created A New user  in Genre points table and updated score  , Wrong Answer")
        }


        if err := db.Where("username = ?",user).Where("genrename = ?" , question.Genrename).Where("quiznumber = ?",question.Quiznumber).First(&personQp).Error ; err == nil {      /// if username ki row exist karti hai to err = nil and we just need to update otherwise create
            fmt.Println("Wrong Answer , but record exists")
        } else {
            // db.Create(&personQp)
            personQp.Username = user
            personQp.Genrename = question.Genrename
            personQp.Quiznumber = question.Quiznumber
            db.Create(&personQp)
            fmt.Println("Created A New user and updated score in QuizPoints Table , Wrong Answer")
        }
        c.Header("access-control-allow-origin", "*")
        c.JSON(200,"Wrong Answer")
        fmt.Println("Wrong Answer")
    }
   
}  

func LeaderBoard(c *gin.Context){
    var people []Person
    var peoples []Person
    if err := db.Find(&people).Error; err != nil {
        c.AbortWithStatus(404)
        fmt.Println(err)
    } else {
        db.Order("totalpoints desc").Find(&peoples)
        c.Header("access-control-allow-origin", "*")
        c.JSON(200,peoples)
        fmt.Println("Leaderboard")
    } 
}

func GenreLeaderBoard(c *gin.Context) {
    gname := c.Params.ByName("genrename")
    var genrelb []Genrepoints
    var glb []Genrepoints

    if err := db.Find(&genrelb).Error ; err != nil {
        c.Header("access-control-allow-origin", "*")
        c.AbortWithError(404,err)
        fmt.Println(err)
    } else {
        db.Order("gpoints desc").Where("genrename = ?", gname ).Find(&glb)
        fmt.Println()
        c.Header("access-control-allow-origin", "*")
        c.JSON(200,glb)
        fmt.Println("GenreLeaderboard")
    }
}

/*func CreatePerson(c *gin.Context) {
    var person Person
    c.BindJSON(&person)
    db.Create(&person)
    c.Header("access-control-allow-origin", "*") // Why am I doing this? Find out. Try running with this line commented
    c.JSON(200, person)
 }*/

func DeleteQuiz(c *gin.Context){
    id := c.Params.ByName("id")
    var quiz Quiz
    db.Where("id = ?",id).First(&quiz)
    fmt.Println(quiz)

    fmt.Println(quiz.Genrename)
    var questions []Question
    if err := db.Where("genrename = ?",quiz.Genrename).Where("quiznumber = ?",quiz.Quiznumber).Find(&questions).Error; err != nil {
        fmt.Println(questions)
        c.Header("access-control-allow-origin", "*")
        c.AbortWithStatus(404)
        fmt.Println("Not Found Any Question")
    } else {
        fmt.Println(questions)
        db.Delete(&quiz)
        db.Where("genrename = ?",quiz.Genrename).Where("quiznumber = ?",quiz.Quiznumber).Delete(Question{})
        c.Header("access-control-allow-origin", "*")
        fmt.Println("Deleted questions of corresponding quiz")
    }
}

func PastQuizzes(c *gin.Context){
    user := c.Params.ByName("username")
    // var genre string
    // var qnum uint
    var quiz []Quizpoints
    if err := db.Where("username = ?", user).Find(&quiz).Error; err != nil {
        c.Header("access-control-allow-origin", "*")
        c.AbortWithError(404,err)
        fmt.Println(err)
        fmt.Println("No quizzes attempted")
     } else {
        c.Header("access-control-allow-origin", "*") // Why am I doing this? Find out. Try running with this line commented
        c.JSON(200, quiz)
        //genre = quiz.Genrename
        //qnum = quiz.Quiznumber    
        //fmt.Println(genre)
        fmt.Println(quiz)
        fmt.Println("All past quizzes returned")
     }   
}


func DeleteQuestion(c *gin.Context){
    id := c.Params.ByName("id")
    var ques Question
    db.Where("id = ?",id).Delete(&ques)
}


func DeletePerson(c *gin.Context) {
   id := c.Params.ByName("id")
   var person Person
   d := db.Where("id = ?", id).Delete(&person)                         // seleect * from ___ where id = __
   fmt.Println(d)
   c.Header("access-control-allow-origin", "*")
   c.JSON(200, gin.H{"id #" + id: "deleted"})
}

func UpdatePerson(c *gin.Context) {
   var person Person
   id := c.Params.ByName("id")
   if err := db.Where("id = ?", id).First(&person).Error; err != nil {
      c.AbortWithStatus(404)
      fmt.Println(err)
   }
   c.BindJSON(&person)
   db.Save(&person)
   c.Header("access-control-allow-origin", "*") // Why am I doing this? Find out. Try running with this line commented
   c.JSON(200, person)
}

func GetPerson(c *gin.Context) {
   id := c.Params.ByName("id")
   var person Person
   if err := db.Where("id = ?", id).First(&person).Error; err != nil {
      c.AbortWithStatus(404)
      fmt.Println(err)
   } else {
      c.Header("access-control-allow-origin", "*") // Why am I doing this? Find out. Try running with this line commented
      c.JSON(200, person)
   }
}

func GetPeople(c *gin.Context) {
   var people []Person
   if err := db.Find(&people).Error; err != nil {
      c.AbortWithStatus(404)
      fmt.Println(err)
   } else {
      c.Header("access-control-allow-origin", "*") // Why am I doing this? Find out. Try running with this line commented
      c.JSON(200, people)
      fmt.Println(people)
   }
}
