/** @type {import('tailwindcss').Config} */
module.exports = {
    content: ['./src/**/*.{html,js,ts,svelte}'],
    darkMode: 'class',
    theme: {
        extend: {},
    },
    daisyui: {
        themes: ['light'],
    },
    plugins: [require('daisyui')],
}

