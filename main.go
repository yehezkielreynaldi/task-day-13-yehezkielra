package main

import (
	"context"
	"day_13/connection"
	"fmt"
	"net/http"
	"strconv"
	"text/template"
	"time"

	"github.com/labstack/echo/v4"
)

type Project struct {
	Id           int
	ProjectName  string
	StartDate    time.Time
	EndDate      time.Time
	Duration     string
	Description  string
	Technologies []string
	Image        string
	Tech1        bool
	Tech2        bool
	Tech3        bool
	Tech4        bool
}

var dataProject = []Project{
	// {
	// 	ProjectName: "Project 1",
	// 	StartDate:   "2023-05-01",
	// 	EndDate:     "2023-06-01",
	// 	Duration:    "1 Bulan",
	// 	Description: "Ini Project 1",
	// 	Tech1:       true,
	// 	Tech2:       true,
	// 	Tech3:       true,
	// 	Tech4:       true,
	// },
	// {
	// 	ProjectName: "Project 2",
	// 	// StartDate:   "2023-05-02",
	// 	// EndDate:     "2023-06-02",
	// 	Duration:    "1 Bulan",
	// 	Description: "Ini Project 2",
	// 	Tech1:       true,
	// 	Tech2:       true,
	// 	Tech3:       true,
	// 	Tech4:       true,
	// },
	// {
	// 	ProjectName: "Project 3",
	// 	StartDate:   "2023-05-03",
	// 	EndDate:     "2023-06-03",
	// 	Duration:    "1 Bulan",
	// 	Description: "Ini Project 3",
	// 	Tech1:       true,
	// 	Tech2:       true,
	// 	Tech3:       true,
	// 	Tech4:       true,
	// },
	// {
	// 	ProjectName: "Project 4",
	// 	StartDate:   "2023-05-04",
	// 	EndDate:     "2023-06-04",
	// 	Duration:    "1 Bulan",
	// 	Description: "Ini Project 4",
	// 	Tech1:       false,
	// 	Tech2:       false,
	// 	Tech3:       true,
	// 	Tech4:       true,
	// },
	// {
	// 	ProjectName: "Project 5",
	// 	StartDate:   "2023-05-05",
	// 	EndDate:     "2023-06-05",
	// 	Duration:    "1 Bulan",
	// 	Description: "Ini Project 5",
	// 	Tech1:       true,
	// 	Tech2:       false,
	// 	Tech3:       true,
	// 	Tech4:       false,
	// },
	// {
	// 	ProjectName: "Project 6",
	// 	StartDate:   "2023-05-06",
	// 	EndDate:     "2023-06-06",
	// 	Duration:    "1 Bulan",
	// 	Description: "Ini Project 6",
	// 	Tech1:       true,
	// 	Tech2:       false,
	// 	Tech3:       true,
	// 	Tech4:       true,
	// },
}

func main() {
	connection.DatabaseConnect()

	e := echo.New()

	// e = echo package
	// GET/POST = run the method
	// "/" = endpoint/routing (ex. localhost:5000'/' | ex. dumbways.id'/lms')
	// helloWorld = function that will run if the routes are opened

	// Serve a static files from "public" directory
	e.Static("/public", "public")

	// Routing

	// GET
	e.GET("/", home)
	e.GET("/contact", contact)
	e.GET("/my-project", myProject)
	e.GET("/project-detail/:id", projectDetail)
	e.GET("/testimonials", testimonials)
	e.GET("/update-project/:id", updateMyProject)

	// POST
	e.POST("/add-project", addProject)
	e.POST("/project-delete/:id", deleteProject)
	e.POST("/update-project/:id", updateProject)

	e.Logger.Fatal(e.Start("localhost:5000"))
}

func home(c echo.Context) error {
	data, _ := connection.Conn.Query(context.Background(), "SELECT id, name, start_date, end_date, description, technologies, image FROM tb_projects")

	var result = []Project{}
	for data.Next() {
		var each = Project{}

		err := data.Scan(&each.Id, &each.ProjectName, &each.StartDate, &each.EndDate, &each.Description, &each.Technologies, &each.Image)
		if err != nil {
			fmt.Println(err.Error())
			return c.JSON(http.StatusInternalServerError, map[string]string{"Message": err.Error()})
		}

		// each.FormatDate = each.PostDate.Format("2 January 2006")
		// each.Author = "Abel Dustin"

		result = append(result, each)
	}

	projects := map[string]interface{}{
		"Projects": result,
	}

	var tmpl, err = template.ParseFiles("views/index.html")

	if err != nil { // null
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": err.Error()})
	}

	return tmpl.Execute(c.Response(), projects)
}

func contact(c echo.Context) error {
	var tmpl, err = template.ParseFiles("views/contact.html")

	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"messsage": err.Error()})
	}

	return tmpl.Execute(c.Response(), nil)
}

func myProject(c echo.Context) error {
	var tmpl, err = template.ParseFiles("views/my-project.html")

	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": err.Error()})
	}

	return tmpl.Execute(c.Response(), nil)
}

func projectDetail(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))

	// data := map[string]interface{}{
	// 	"Id":      id,
	// 	"Title":   "Dumbways Web App",
	// 	"Content": `Lorem ipsum dolor sit amet consectetur adipisicing elit. Rerum nam excepturi reprehenderit molestias vero laboriosam, accusamus aliquam? Voluptatem, libero consectetur quaerat aspernatur porro quidem error facere reprehenderit omnis earum nisi quos aperiam soluta vel tempora dignissimos possimus facilis quas, animi eaque nostrum suscipit perferendis optio ullam? Praesentium excepturi animi eius illum autem voluptates labore. Libero excepturi nisi ipsam veritatis est voluptatibus voluptates recusandae sapiente dolore distinctio! Cumque asperiores corporis necessitatibus, quisquam neque adipisci. Itaque, natus harum sint eum nesciunt ea ipsa perferendis porro soluta magni, corporis asperiores accusamus sed minus? Laudantium aperiam rem beatae voluptatum ipsum ipsam at dignissimos nobis. <br /> Lorem ipsum dolor sit amet consectetur adipisicing elit. Magni nam animi officiis ipsum? Eligendi voluptas dolorem, ab consequatur, neque magni veniam, modi quaerat labore quo ut dolorum velit a voluptates sint dolores dolor! Similique dolores unde adipisci neque iure exercitationem, distinctio sed debitis officiis nulla vero! Deleniti veniam quae, veritatis accusantium vero dicta, nihil modi atque voluptatem dolor rem dolorum dolore eum sequi harum nesciunt repellat quidem vitae architecto! Molestiae ipsa reprehenderit nam deserunt assumenda a magni, harum pariatur at exercitationem rem officiis ipsum repellendus nulla in laudantium rerum delectus natus facilis. Itaque nesciunt eveniet debitis consectetur veniam repellat modi! <br /> Lorem ipsum dolor sit amet consectetur adipisicing elit. Soluta quia eaque quae voluptatem nisi odit obcaecati rem vero ullam in porro maiores aliquam alias enim consequatur vitae tempora, nesciunt modi beatae animi. Dignissimos fugiat corrupti amet mollitia, et esse rem voluptas fuga incidunt dolores iure consectetur repudiandae placeat, minima adipisci praesentium nobis debitis minus distinctio atque. Aliquid mollitia totam nemo natus. Sunt et culpa blanditiis, commodi tempore eligendi itaque eius aliquam? Sequi incidunt molestiae odio cupiditate dicta voluptates a et commodi, facere nihil ratione ipsa sit quibusdam, ipsum autem mollitia? Quae fugit laboriosam cum numquam perferendis aperiam laudantium et vitae. <br /> Lorem ipsum dolor sit amet consectetur, adipisicing elit. Ipsam animi doloremque iste, ab iure vero praesentium repellat harum quisquam amet id dolorum unde dolorem magnam laborum modi. Earum voluptatibus, tempora minima vel culpa minus perferendis sapiente nostrum harum, inventore quaerat voluptates, explicabo obcaecati repudiandae possimus unde ullam sequi suscipit accusamus ea! Mollitia odio harum omnis nam porro aut corporis nisi sit nobis nulla, tempore explicabo animi non corrupti ipsam, libero ratione cum minus esse reiciendis! Dolore laboriosam, ab provident fuga sapiente praesentium natus aspernatur impedit excepturi ea quaerat sunt quasi voluptates ipsum veritatis architecto. Porro quas quos ratione eligendi rerum.`,
	// }

	var ProjectDetail = Project{}

	for i, data := range dataProject {
		if id == i {
			ProjectDetail = Project{
				ProjectName: data.ProjectName,
				StartDate:   data.StartDate,
				EndDate:     data.EndDate,
				Duration:    data.Duration,
				Description: data.Description,
				Tech1:       data.Tech1,
				Tech2:       data.Tech2,
				Tech3:       data.Tech3,
				Tech4:       data.Tech4,
			}
		}
	}

	data := map[string]interface{}{
		"Project": ProjectDetail,
	}

	var tmpl, err = template.ParseFiles("views/project-detail.html")

	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": err.Error()})
	}

	return tmpl.Execute(c.Response(), data)
}

func testimonials(c echo.Context) error {
	var tmpl, err = template.ParseFiles("views/testimonials.html")

	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": err.Error()})
	}

	return tmpl.Execute(c.Response(), nil)
}

func updateMyProject(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))

	var ProjectDetail = Project{}

	for i, data := range dataProject {
		if id == i {
			ProjectDetail = Project{
				Id:          id,
				ProjectName: data.ProjectName,
				StartDate:   data.StartDate,
				EndDate:     data.EndDate,
				Duration:    data.Duration,
				Description: data.Description,
				Tech1:       data.Tech1,
				Tech2:       data.Tech2,
				Tech3:       data.Tech3,
				Tech4:       data.Tech4,
			}
		}
	}

	data := map[string]interface{}{
		"Project": ProjectDetail,
	}

	var tmpl, err = template.ParseFiles("views/update-project.html")

	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": err.Error()})
	}

	return tmpl.Execute(c.Response(), data)
}

func addProject(c echo.Context) error {
	projectName := c.FormValue("projectName")
	startDate := c.FormValue("startDate")
	endDate := c.FormValue("endDate")
	duration := hitungDurasi(startDate, endDate)
	description := c.FormValue("descriptionProject")
	tech1 := c.FormValue("tech1")
	tech2 := c.FormValue("tech2")
	tech3 := c.FormValue("tech3")
	tech4 := c.FormValue("tech4")

	println("Project Name : " + projectName)
	println("Start Date : " + startDate)
	println("End Date : " + endDate)
	println("Description : " + description)
	println("Technologies : " + tech1)
	println("Technologies : " + tech2)
	println("Technologies : " + tech3)
	println("Technologies : " + tech4)

	var newProject = Project{
		ProjectName: projectName,
		// StartDate:   startDate,
		// EndDate:     endDate,
		Duration:    duration,
		Description: description,
		Tech1:       (tech1 == "tech1"),
		Tech2:       (tech2 == "tech2"),
		Tech3:       (tech3 == "tech3"),
		Tech4:       (tech4 == "tech4"),
	}

	dataProject = append(dataProject, newProject)

	fmt.Println(dataProject)

	return c.Redirect(http.StatusMovedPermanently, "/")
}

func deleteProject(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))

	fmt.Println("Index : ", id)

	dataProject = append(dataProject[:id], dataProject[id+1:]...)

	return c.Redirect(http.StatusMovedPermanently, "/")
}

func updateProject(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))

	fmt.Println("Index :", id)

	projectName := c.FormValue("projectName")
	startDate := c.FormValue("startDate")
	endDate := c.FormValue("endDate")
	duration := hitungDurasi(startDate, endDate)
	description := c.FormValue("descriptionProject")
	tech1 := c.FormValue("tech1")
	tech2 := c.FormValue("tech2")
	tech3 := c.FormValue("tech3")
	tech4 := c.FormValue("tech4")

	println("Project Name : " + projectName)
	println("Start Date : " + startDate)
	println("End Date : " + endDate)
	println("Description : " + description)
	println("Technologies : " + tech1)
	println("Technologies : " + tech2)
	println("Technologies : " + tech3)
	println("Technologies : " + tech4)

	var updateProject = Project{
		ProjectName: projectName,
		// StartDate:   startDate,
		// EndDate:     endDate,
		Duration:    duration,
		Description: description,
		Tech1:       (tech1 == "tech1"),
		Tech2:       (tech2 == "tech2"),
		Tech3:       (tech3 == "tech3"),
		Tech4:       (tech4 == "tech4"),
	}

	dataProject[id] = updateProject

	return c.Redirect(http.StatusMovedPermanently, "/")
}

func hitungDurasi(startDate, endDate string) string {
	startTime, _ := time.Parse("2006-01-02", startDate)
	endTime, _ := time.Parse("2006-01-02", endDate)

	durationTime := int(endTime.Sub(startTime).Hours())
	durationDays := durationTime / 24
	durationWeeks := durationDays / 7
	durationMonths := durationWeeks / 4
	durationYears := durationMonths / 12

	var duration string

	if durationYears > 1 {
		duration = strconv.Itoa(durationYears) + " Tahun"
	} else if durationYears > 0 {
		duration = strconv.Itoa(durationYears) + "Tahun"
	} else {
		if durationMonths > 1 {
			duration = strconv.Itoa(durationMonths) + " Bulan"
		} else if durationMonths > 0 {
			duration = strconv.Itoa(durationMonths) + " Bulan"
		} else {
			if durationWeeks > 1 {
				duration = strconv.Itoa(durationWeeks) + "Minggu"
			} else if durationWeeks > 0 {
				duration = strconv.Itoa(durationWeeks) + "Minggu"
			} else {
				if durationDays > 1 {
					duration = strconv.Itoa(durationDays) + " Hari"
				} else {
					duration = strconv.Itoa(durationDays) + " Hari"
				}
			}
		}
	}

	return duration
}
