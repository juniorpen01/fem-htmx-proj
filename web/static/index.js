document.addEventListener("htmx:beforeSwap", (e) => {
  if (e.detail.xhr.status == 409) {
    e.detail.shouldSwap = true;
    e.detail.isError = false; // idk if this works
  }
});
