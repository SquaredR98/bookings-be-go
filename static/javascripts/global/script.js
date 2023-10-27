const menuContainer = document.getElementById("menu-container");
const menuItems = document.getElementById("menu-items");

document.getElementById("menu-bar").addEventListener("click", () => {
  if (menuContainer.classList.contains("hidden")) {
    menuContainer.classList.add(
      "fixed",
      "h-screen",
      "w-3/4",
      "right-0",
      "top-0",
      "flex",
      "z-[9999]"
    );
    menuItems.classList.add(
      "flex",
      "flex-col",
      "justify-center",
      "items-center",
    );
    menuItems.querySelectorAll("*").forEach((el) => {
      el.classList.add("text-xl", "text-center");
    });
    document
      .getElementById("close-btn")
      .classList.add("flex", "justify-end", "m-4");
    menuContainer.classList.remove("hidden");
  } else {
    menuContainer.classList.add("hidden");
  }
});

document.getElementById("close-btn").addEventListener("click", () => {
  if (!menuContainer.classList.contains("hidden")) {
    menuContainer.classList.add("hidden");
  }
});

const suitesLink = document.getElementById("suites-link");
const dropdownIcon = document.getElementById("dropdown-icon");
const navbarDropdown = document.getElementById("navbar-dropdown");

function showDropdown() {
  navbarDropdown.classList.add("z-[9999]");
  navbarDropdown.classList.remove("hidden");
}
function hideDropdown() {
  navbarDropdown.classList.add("hidden");
}

suitesLink.addEventListener("mouseover", showDropdown);
navbarDropdown.addEventListener("mouseleave", hideDropdown);

document.addEventListener("click", function (event) {
  if (event.target !== suitesLink) {
    hideDropdown();
  }
});

function notify(msg, msgType) {
  Swal.fire({
    toast: true,
    position: 'top-end',
    icon: msgType,
    title: msg,
    showConfirmButton: false,
    timer: 1500
  })
}

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
    Toast.fire({});
  };

  let success = function (c) {
    const { msg = "", title = "", footer = "" } = c;

    Swal.fire({
      icon: "success",
      title: title,
      text: msg,
      footer: footer,
    });
  };

  let error = function (c) {
    const { msg = "", title = "", footer = "" } = c;

    Swal.fire({
      icon: "error",
      title: title,
      text: msg,
      footer: footer,
    });
  };

  async function custom(c) {
    const { msg = "", title = "" } = c;

    const { value: result } = await Swal.fire({
      title: title,
      html: msg,
      backdrop: false,
      focusConfirm: false,
      showCancelButton: true,
      willOpen: () => c.willOpen && c.willOpen(),
      didOpen: () => c.didOpen && c.didOpen(),
      preConfirm: () => {
        return [
          document.getElementById("startDate").value,
          document.getElementById("endDate").value,
        ];
      },
    });

    if (result) {
      console.log(result)
      if (result.dismiss !== Swal.DismissReason.cancel) {
        if (result.value !== "") {
          if (c.callback !== undefined) {
            c.callback(result);
          }
        } else {
          c.callback(false);
        }
      } else {
        c.callback(false);
      }
    }
  }

  return {
    toast: toast,
    success: success,
    error: error,
    custom: custom,
  };
}
