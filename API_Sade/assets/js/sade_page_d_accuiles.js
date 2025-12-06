// Subtle ambient interactions for the Sade artiste page

(function () {
  const shell = document.querySelector('.page-shell');
  if (!shell) return;

  // Soft parallax / glow on pointer move
  shell.addEventListener('pointermove', (e) => {
    const rect = shell.getBoundingClientRect();
    const x = (e.clientX - rect.left) / rect.width - 0.5;
    const y = (e.clientY - rect.top) / rect.height - 0.5;

    shell.style.transform = `translateY(${y * -4}px)`;
    shell.style.boxShadow = `0 18px 45px rgba(0,0,0,0.75), ${x * 8}px ${y * 8}px 40px rgba(195,162,122,0.16)`;
  });

  shell.addEventListener('pointerleave', () => {
    shell.style.transform = '';
    shell.style.boxShadow = '';
  });
})();
