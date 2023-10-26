$("#menu-bar").on("click", function () {
  if ($("#menu-container").hasClass("hidden")) {
    $("#menu-container").addClass(
      "fixed flex-col h-screen w-3/4 right-0 top-0 flex"
    );
    $("#menu-items").addClass(
      "flex flex-col justify-center items-center h-full"
    );
    $("#menu-items *").addClass("text-xl w-full text-center py-2");
    $("#close-btn").addClass("flex justify-end m-4");
    $("#menu-container").removeClass("hidden");
  } else {
    $("#menu-container").addClass("hidden");
  }
});

$("#close-btn").on("click", function () {
  if (!$("#menu-container").hasClass("hidden")) {
    $("#menu-container").addClass("hidden");
  } else {
  }
});

const suitesLink = document.getElementById('suites-link');
const dropdownIcon = document.getElementById('dropdown-icon')
const navbarDropdown = document.getElementById('navbar-dropdown')

// Function to show the dropdown
function showDropdown() {
  navbarDropdown.classList.remove('hidden');
}

// Function to hide the dropdown
function hideDropdown() {
  navbarDropdown.classList.add("hidden");
}

// Event listener to show the dropdown when the link is hovered over
suitesLink.addEventListener("mouseover", showDropdown);

// Event listener to hide the dropdown when the mouse leaves the dropdown area
navbarDropdown.addEventListener("mouseleave", hideDropdown);

// Event listener to hide the dropdown when clicking outside of it
document.addEventListener("click", function (event) {
  if (event.target !== suitesLink) {
    hideDropdown();
  }
});

function Prompt() {
  let toast = function (data) {
    const { message = "", icon = "success", position = "top-end" } = data;
    const Toast = Swal.mixin({
      toast: true,
      position,
      title: message,
      showConfirmationButton: false,
      icon,
      timer: 3000,
      timerProgressBar: true,
      showConfirmButton: false,
      didOpen: (toast) => {
        toast.addEventListener("mouseenter", Swal.stopTimer);
        toast.addEventListener("mouseleave", Swal.resumeTimer);
      },
    });
    Toast.fire({})
  };

  return { toast }
}

const prompt = Prompt();
prompt.toast({ message: "Hello Guys" })
