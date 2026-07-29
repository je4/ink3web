
(() => {
  'use strict'

  /* theme switch */
  // theme switch fucntions are located in the head.gohtml file, because they are needed before the page is loaded to avoid flickering
  
  // redo the init to change the nav button to the corect state
  initTheme();
  
  // listeen to theme changes
  document.addEventListener("click", e => {
    const btn = e.target.closest("[data-bs-theme-value]");
    if (!btn) return;
    console.log('theme:'+btn.dataset.bsThemeValue);
    setTheme(btn.dataset.bsThemeValue);
  })

  /* read. more */
  document.querySelectorAll(".textclamp").forEach(el => {
    const readMore = el.querySelector(".readmore");
    if (readMore) {
      readMore.hidden = el.scrollHeight <= el.clientHeight;
    }
  });

  document.addEventListener("click", e => {
    const btn = e.target.closest(".readmore");
    if (!btn) return;

    e.preventDefault();

    const container = btn.closest(".textclamp");
    if (!container) return;

    container.classList.remove("textclamp");
    btn.hidden = true; // optional: "Read more" ausblenden
  });
})()

function historyOrLink(link) {
    if (window.history.length > 1) {
        window.history.back();
    } else {
        window.location.href = link;
    }
}