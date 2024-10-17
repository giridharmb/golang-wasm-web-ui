package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/maxence-charriere/go-app/v10/pkg/app"
)

type FormData struct {
	Name      string   `json:"name"`
	Email     string   `json:"email"`
	Age       int      `json:"age"`
	Gender    string   `json:"gender"`
	Interests []string `json:"interests"`
	Comments  string   `json:"comments"`
}

type ComplexForm struct {
	app.Compo
	formData          FormData
	selectedInterests map[string]bool
	submitStatus      string
}

func (f *ComplexForm) Render() app.UI {
	return app.Div().
		Class("min-h-screen bg-gradient-to-r from-purple-400 via-pink-500 to-red-500 flex items-center justify-center p-4").
		Body(
			app.Div().
				Class("bg-white rounded-lg shadow-2xl p-8 max-w-md w-full").
				Body(
					app.H1().
						Class("text-3xl font-bold mb-6 text-center text-gray-800").
						Text("Awesome Web Form"),
					app.Form().
						Class("space-y-6").
						OnSubmit(f.handleSubmit).
						Body(
							f.renderInput("name", "Name", "text", f.formData.Name),
							f.renderInput("email", "Email", "email", f.formData.Email),
							f.renderInput("age", "Age", "number", strconv.Itoa(f.formData.Age)),
							f.renderSelect(),
							f.renderInterests(),
							f.renderTextarea(),
							app.Button().
								Type("submit").
								Class("w-full bg-gradient-to-r from-purple-500 to-pink-500 text-white font-bold py-3 px-4 rounded-full hover:from-pink-500 hover:to-purple-500 transition duration-300 ease-in-out transform hover:-translate-y-1 hover:scale-105 focus:outline-none focus:ring-2 focus:ring-purple-600 focus:ring-opacity-50").
								Text("Submit"),
						),
					app.P().
						Class("mt-4 text-center text-gray-600").
						Text(f.submitStatus),
				),
		)
}

func (f *ComplexForm) renderInput(id, label, inputType, value string) app.UI {
	return app.Div().Body(
		app.Label().
			For(id).
			Class("block text-sm font-medium text-gray-700 mb-1").
			Text(label),
		app.Input().
			ID(id).
			Type(inputType).
			Class("mt-1 block w-full rounded-md border-gray-300 shadow-sm focus:border-purple-500 focus:ring focus:ring-purple-500 focus:ring-opacity-50 text-center px-4 py-2").
			Value(value).
			OnInput(func(ctx app.Context, e app.Event) {
				switch id {
				case "name":
					f.formData.Name = e.Get("target").Get("value").String()
				case "email":
					f.formData.Email = e.Get("target").Get("value").String()
				case "age":
					age, _ := strconv.Atoi(e.Get("target").Get("value").String())
					f.formData.Age = age
				}
			}),
	)
}

func (f *ComplexForm) renderInterests() app.UI {
	interests := []string{"Sports", "Music", "Reading", "Travel", "Technology"}
	return app.Div().Body(
		app.Label().
			Class("block text-sm font-medium text-gray-700 mb-2").
			Text("Interests"),
		app.Div().
			Class("space-y-2").
			Body(
				app.Range(interests).Slice(func(i int) app.UI {
					interest := interests[i]
					isChecked := f.selectedInterests[interest]
					return app.Div().
						Class("flex items-center").
						Body(
							app.Input().
								ID("interest-"+interest).
								Type("checkbox").
								Class("h-4 w-4 text-purple-600 focus:ring-purple-500 border-gray-300 rounded").
								Checked(isChecked).
								OnChange(func(ctx app.Context, e app.Event) {
									f.selectedInterests[interest] = e.Get("target").Get("checked").Bool()
									f.updateInterests()
								}),
							app.Label().
								For("interest-"+interest).
								Class("ml-2 block text-sm text-gray-700").
								Text(interest),
						)
				}),
			),
	)
}

func (f *ComplexForm) updateInterests() {
	f.formData.Interests = []string{}
	for interest, selected := range f.selectedInterests {
		if selected {
			f.formData.Interests = append(f.formData.Interests, interest)
		}
	}
}

// Ensure all other methods are properly defined

func (f *ComplexForm) onInputName(ctx app.Context, e app.Event) {
	f.formData.Name = e.Get("target").Get("value").String()
}

func (f *ComplexForm) onInputEmail(ctx app.Context, e app.Event) {
	f.formData.Email = e.Get("target").Get("value").String()
}

func (f *ComplexForm) onInputAge(ctx app.Context, e app.Event) {
	age, _ := strconv.Atoi(e.Get("target").Get("value").String())
	f.formData.Age = age
}

func (f *ComplexForm) onChangeGender(ctx app.Context, e app.Event) {
	f.formData.Gender = e.Get("target").Get("value").String()
}

func (f *ComplexForm) onInputComments(ctx app.Context, e app.Event) {
	f.formData.Comments = e.Get("target").Get("value").String()
}

func (f *ComplexForm) renderTextarea() app.UI {
	return app.Div().Body(
		app.Label().
			For("comments").
			Class("block text-sm font-medium text-gray-700 mb-1").
			Text("Comments"),
		app.Textarea().
			ID("comments").
			Class("mt-1 block w-full rounded-md border-gray-300 shadow-sm focus:border-purple-500 focus:ring focus:ring-purple-500 focus:ring-opacity-50 px-4 py-2").
			Rows(4).
			Text(f.formData.Comments).
			OnInput(f.onInputComments),
	)
}

func (f *ComplexForm) renderSelect() app.UI {
	return app.Div().Body(
		app.Label().
			For("gender").
			Class("block text-sm font-medium text-gray-700 mb-1").
			Text("Gender"),
		app.Select().
			ID("gender").
			Class("mt-1 block w-full rounded-md border-gray-300 shadow-sm focus:border-purple-500 focus:ring focus:ring-purple-500 focus:ring-opacity-50").
			OnChange(f.onChangeGender).
			Body(
				app.Option().Value("").Text("Select...").Selected(f.formData.Gender == ""),
				app.Option().Value("male").Text("Male").Selected(f.formData.Gender == "male"),
				app.Option().Value("female").Text("Female").Selected(f.formData.Gender == "female"),
				app.Option().Value("other").Text("Other").Selected(f.formData.Gender == "other"),
			),
	)
}

// Other methods (onInputName, onInputEmail, onInputAge, onChangeGender, onInputComments, updateInterests) remain the same

func (f *ComplexForm) handleSubmit(ctx app.Context, e app.Event) {
	e.PreventDefault()

	jsonData, err := json.Marshal(f.formData)
	if err != nil {
		f.submitStatus = "Error preparing data"
		return
	}

	go func() {
		resp, err := http.Post("http://localhost:8080/submit", "application/json", bytes.NewBuffer(jsonData))
		if err != nil {
			f.submitStatus = "Error submitting form"
			ctx.Dispatch(func(ctx app.Context) {})
			return
		}
		defer resp.Body.Close()

		var result map[string]interface{}
		json.NewDecoder(resp.Body).Decode(&result)

		f.submitStatus = fmt.Sprintf("Form submitted successfully. Server response: %v", result["message"])
		ctx.Dispatch(func(ctx app.Context) {})
	}()
}

func main() {
	// Set up the route for the web app
	app.Route("/", func() app.Composer {
		return &ComplexForm{
			selectedInterests: make(map[string]bool),
		}
	})

	// Run the app when in the browser
	app.RunWhenOnBrowser()

	// Serve static files
	http.Handle("/web/", http.StripPrefix("/web/", http.FileServer(http.Dir("web"))))

	// Serve index.html for the root path
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "web/index.html")
	})

	// Set up the handler for the web app
	handler := &app.Handler{
		Name:        "Awesome Complex Form",
		Description: "An awesome complex web form using go-app",
		Styles: []string{
			"https://cdn.jsdelivr.net/npm/tailwindcss@2.2.19/dist/tailwind.min.css",
		},
	}

	fmt.Println("Server running on http://localhost:8000")
	if err := http.ListenAndServe(":8000", handler); err != nil {
		log.Fatal(err)
	}
}
