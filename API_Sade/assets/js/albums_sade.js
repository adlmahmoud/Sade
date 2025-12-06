// Light interaction for albums table: highlight row on focus / keyboard nav

(function () {
  const table = document.querySelector('#albums-table table');
  if (!table) return;

  table.querySelectorAll('tbody tr').forEach((row) => {
    row.addEventListener('click', () => {
      row.classList.add('is-active');
      setTimeout(() => row.classList.remove('is-active'), 260);
    });
  });
})();
