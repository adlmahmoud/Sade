// Ambient highlighting of active track row

(function () {
  const table = document.querySelector('#tracks-table table');
  if (!table) return;

  table.querySelectorAll('tbody tr').forEach((row) => {
    row.addEventListener('mouseenter', () => {
      row.classList.add('is-hovered');
    });
    row.addEventListener('mouseleave', () => {
      row.classList.remove('is-hovered');
    });
  });
})();
