const addErrorBorderAndRemove = (elem) => {
  elem.classList.add("border-red-600")
  setTimeout(() => {
    elem.classList.remove("border-red-600")
  }, 3000)
}

const emptyField = (fieldValue) => {
  return fieldValue === ''
}

const formValidation = () => {
  const form = document.getElementById('reservation-form')
  const userName = form.querySelector('input[name=name]')
  const userEmail = form.querySelector('input[name=email]')
  const userContact = form.querySelector('input[name=mobile]')
  const button = form.querySelector('button')
  // button.addEventListener('click', (event) => {
  //   [userName, userEmail, userContact].forEach((el, idx) => {
  //     if(emptyField(el.value)) addErrorBorderAndRemove(el) && event.preventDefault()
  //   })
  // })
}
window.onload = formValidation
