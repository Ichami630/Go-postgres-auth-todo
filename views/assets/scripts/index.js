const menuBar = document.querySelector(".hide-menu");
const topBarText =  document.querySelector(".top-menu-text");
const asideContent =  document.querySelector(".aside-content");
const menuAside = document.querySelector(".menu-aside")
const dashboardContent = document.querySelector(".dashboard-content")

const screenWidth = window.screen.availWidth;

menuBar.addEventListener('click',()=>{
    if(topBarText.style.display === "none"){
        console.log(screenWidth)
        if(screenWidth < 574){
            displayBlock(topBarText)
            displayBlock(asideContent)
            menuAside.style.background = "#d3d1d1"
            menuAside.style.zIndex = "50"
            menuAside.style.gridColumn = "span 12"
            dashboardContent.style.gridColumn = "span 12"
        }else{
            displayBlock(topBarText)
            displayBlock(asideContent)
            menuAside.style.background = "#d3d1d1"
            menuAside.style.gridColumn = "span 3"
            dashboardContent.style.gridColumn = "span 9"
        }
    }else{
        displayNone(topBarText)
        displayNone(asideContent)
        menuAside.style.background = "#FFFF"
        menuAside.style.gridColumn = "span 1"
        dashboardContent.style.gridColumn = "span 11"

    }
})

function displayNone(target){
    target.style.display = "none"
}
function displayBlock(target){
    target.style.display = "block"
}