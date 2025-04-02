
document.addEventListener("DOMContentLoaded", function () {
    if (document.getElementById("signup-form")) {
        signup();
    }
    if (document.getElementById("login-form")) {
        login();
    }
});
//signup
function signup(){
    document.getElementById("signup-form").addEventListener("submit", async function(event) {
        event.preventDefault(); // Prevent form from reloading page
    
        const email = document.getElementById("email").value;
        const password = document.getElementById("password").value;
        const cpassword = document.getElementById("cpassword").value;
        const csrf_token = document.getElementById("csrf_token").value;
        const message = document.getElementById("message");
    
        // Send data to backend using Fetch API
        const response = await fetch("/signup", {
            method: "POST",
            headers: { 
                "Content-Type": "application/json",
                "X-CSRF-Token": csrf_token, 
            },
            body: JSON.stringify({ email, password, cpassword })
        });
    
        const result = await response.json(); // Parse JSON response
        if (response.ok) {
            message.style.color = "green";
            message.innerText = result.success;
        } else {
            message.style.color = "red";
            message.innerText = result.error;
        }
    
        // Display success or error message dynamically
        setTimeout(()=>{
            message.innerHTML = ""
        },5000)
    });
}

//login
function login(){
    document.getElementById("login-form").addEventListener("submit",async function(event) {
        event.preventDefault()
        //send data to backend fetch api
        const email = document.getElementById("email").value;
        const password = document.getElementById("password").value;

        const response = await fetch("/login", {
            method: "POST",
            headers: { "Content-Type": "application/json" },
            body: JSON.stringify({email,password})
        })
        const result = await response.json(); // Parse JSON response
        if (response.ok) {
            message.style.color = "green";
            message.innerText = result.success;
        } else {
            message.style.color = "red";
            message.innerText = result.error;
        }

    });
}
