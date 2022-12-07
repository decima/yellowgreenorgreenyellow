const defaultColors = require('daisyui/src/colors/themes')


let colors = {
    ...defaultColors["[data-theme=light]"],
    ...{
        primary: "yellowgreen",
        secondary: "greenyellow",
    }
}

/** @type {import('tailwindcss').Config} */
module.exports = {
    content: [
        "./index.html",
        "./src/**/*.{svelte,js,ts,jsx,tsx}",

    ],
    theme: {
        extend: {},
    },
    plugins: [require("daisyui"), require('@tailwindcss/typography')],
    daisyui: {
        themes: [
            {
                mytheme: {
                    "primary": "greenyellow",
                    "secondary": "yellowgreen",
                    "accent": "rebeccapurple",
                    "neutral": "#19242E",
                    "base-100": "ghostwhite",
                    "info": "#6AD6F1",
                    "success": "#289F56",
                    "warning": "#F1CB22",
                    "error": "#F15581",
                },
            },
        ],
    }
}