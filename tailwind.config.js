/** @type {import('tailwindcss').Config} */
module.exports = {
  darkMode: 'class',
  content: ["./templates/**/*.{html,js}"],
  theme: {
    container: {
      center: true,
    },
    extend: {
      fontFamily: {
        sans: ["Iosevka Aile Iaso", "sans-serif"],
        mono: ["Iosevka Curly Iaso", "monospace"],
        serif: ["Iosevka Etoile Iaso", "serif"],
      },
      typography: (theme) => ({
        DEFAULT: {
          css: {
            color: '#3a3737',
            h2: {
              color: '#3a3737',
            },
            pre: {
              backgroundColor: 'black',
              color: 'silver',
            },
            'ul > li::marker': {
              color: '#3a3737',
            },
            a: {
              color: '#6D214F'
            }
          },
        },
        dark: {
          css: {
            color: '#fbf1c7',
            h2: {
              color: '#fbf1c7',
            },
            pre: {
              backgroundColor: '#32302f',
              color: '#fbf1c7',
            },
            'ul > li::marker': {
              color: '#fbf1c7',
            },
            a: {
              color: '#D6A2E8'
            },
          },
        },
      }),
    },
  },
  plugins: [
    require('@tailwindcss/typography'),
  ],
}
