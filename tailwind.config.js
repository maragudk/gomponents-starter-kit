const colors = require('tailwindcss/colors')
const defaultTheme = require('tailwindcss/defaultTheme')

/** @type {import('tailwindcss').Config} */
module.exports = {
  content: [
    "./html/**/*.go",
  ],
  theme: {
    extend: {
      fontFamily: {
        mono: [...defaultTheme.fontFamily.mono],
        sans: [...defaultTheme.fontFamily.sans],
        serif: [...defaultTheme.fontFamily.serif],
      },
    }
  },
  plugins: [
    require('@tailwindcss/typography'),
  ],
}
