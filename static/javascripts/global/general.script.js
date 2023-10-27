const button = document
  .getElementById("check-availability-button")
  .addEventListener("click", () => {
    const html = `<form id="availability-form" class="w-11/12 mx-auto">
    <div
      class="flex flex-col w-full lg:w-full md:flex-row mx-auto"
      id="range-picker"
    >
      <div class="md:w-1/2 md:mr-2">
        <label for="startDate" hidden class="text-xl font-bold">Start Date</label>
        <input
          required
          name="startDate"
          id="startDate"
          placeholder="Arrival"
          class="rounded-md my-2 outline-none border p-2 w-full"
        />
      </div>
      <div class="md:w-1/2 md:ml-2">
        <label for="endDate" hidden class="text-xl font-bold">End Date</label>
        <input
          required
          id="endDate"
          placeholder="Departure"
          name="endDate"
          class="rounded-md my-2 outline-none border p-2 w-full"
        />
      </div>
    </div>
  </form>`;

    Prompt().custom({
      msg: html,
      title: "Choose Your Dates",
      willOpen: () => {
        const elem = document.getElementById("range-picker");
        const rp = new DateRangePicker(elem, {
          format: "yyyy-mm-dd",
          showOnFocus: false,
        });
      },
      didOpen: () => {
        document.getElementById("startDate").removeAttribute("disabled");
        document.getElementById("endDate").removeAttribute("disabled");
      },
      callback: (result) => {
        let data = document.getElementById("availability-form");
        let formData = new FormData(data);
        formData.append("csrf_token", "{{.CSRFToken}}");
        fetch("/availability", {
          method: "POST",
          body: formData,
        })
          .then((response) => response.json())
          .then((data) => {
            console.log(data);
          })
          .catch((err) => err);
      },
    });
  });
