## Intro

This is my portfolio website. It is made using Go, html/templates from the Go standard library and CSS for styling. It is hosted on AWS using Docker. 

## Use Case

- You can use this to make your own website easily. 
- I use this general format to make mine and my client's website as the design and easy to use a modular. 

## How is it stuctured

### cmd/web

all the server code is in this folder. 

- the routes are defined in to routes.go file. 
    - I use "github.com/bmizerany/pat" for the routing although I can use the standard library now for this as well I initially wrote this before the Go 1.22 release. 

    - I use "github.com/justinas/alice" for chaining my middleware, just a quality of life thing. 
- templates.go as you can see I use the html/templates from the standard lib for ganarating my html. in this file I define my struct for the templates and also create a function called newTemplateCache() this creates a cache for all the templates in this project because I have contact, home and projects I have those in the cache which can be called through the app.render function which renders the html. Because the html is not being read from disk at the moment of request this makes te website faster to load. This is generally a good stategy for smaller websites

- The rest for the files are pretty standard and self explanatory

### pkg

Package has the components that can be taken out with ease, if your portfolio only has the static html, and no forms you can take the forms section out even from templates struct safely.

and if your website does not require a contact form then it is the same thing you cna take the mailing pkg out. 

### ui/html and ui/static

these are self explanatory

### how do I use mailing

I personally use it such that I email myself, because then even if the user fails to give the right email I always get the email. 

## templates

- All the templates are inside the ui/html directory. 

- The base.layout.tmpl page defines the standard layout that will be repeated with every page, it includes boilerpate html, header, and it gives room for body and more or less a footer. 

- the *.page.tmpl defines the content of the page, whatever that may be, in the page you have to define the "title" and "body"

- The partial could be almost thing you'd want and can be used in any part (layout, page, partial), even though I only use it as a footer. 

- another use case for the partial would be to use it somewhat like react make components like botton ex.adding a specific styling to them or adding a js script to add specific functionality to it. 

## Conclusion 

I hope this helps, I tried to explain it as much detail as I could. Use this in any way you like 
