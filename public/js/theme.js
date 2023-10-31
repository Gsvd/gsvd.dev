function toggleTheme() {
    if (localStorage.theme === 'dark') {
        document.getElementById('theme-indicator').src = '/images/sun.svg';
        document.documentElement.classList.remove('dark');
        localStorage.setItem('theme', 'light');
    } else {
        document.getElementById('theme-indicator').src = '/images/moon.svg';
        document.documentElement.classList.add('dark');
        localStorage.setItem('theme', 'dark');
    }
}

if (localStorage.theme === 'dark' || (!('theme' in localStorage) && window.matchMedia('(prefers-color-scheme: dark)').matches)) {
    document.getElementById('theme-indicator').src = '/images/moon.svg';
    document.documentElement.classList.add('dark');
    localStorage.setItem('theme', 'dark');
} else {
    document.getElementById('theme-indicator').src = '/images/sun.svg';
    document.documentElement.classList.remove('dark');
    localStorage.setItem('theme', 'light');
}