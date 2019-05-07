#
# This is a Shiny web application. You can run the application by clicking
# the 'Run App' button above.
#
# Find out more about building applications with Shiny here:
#
#    http://shiny.rstudio.com/
#

library(shiny)
library(shinythemes)

# 

ui <- navbarPage( HTML("<strong>DLab</strong><small>cloud</small>"),
                  tabPanel("Species",
                           sidebarLayout(
                               sidebarPanel(
                                   radioButtons("plotType", "Plot type",
                                                c("Scatter"="p", "Line"="l")
                                   )
                               ),
                               mainPanel(
                                   plotOutput("plot")
                               )
                           )),
                  tabPanel("Table",
                           DT::dataTableOutput("table")
                  ),
                  theme = shinytheme("cyborg")
)


# Define server logic required to draw a histogram
server <- function(input, output) {
    set.seed(122)
    histdata <- rnorm(500)
    
    output$plot <- renderPlot({
        plot(cars, type=input$plotType)
    })
    
    output$table <- DT::renderDataTable({
        DT::datatable(cars, style="bootstrap")
    })
}

# Run the application 
shinyApp(ui = ui, server = server)
