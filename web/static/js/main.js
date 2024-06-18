const allowedStatus = [422, 401, 500];

document.body.addEventListener("htmx:beforeSwap", function (e) {
  if (allowedStatus.includes(e.detail.xhr.status)) {
    e.detail.shouldSwap = true;
    e.detail.isError = true;
  }
});
