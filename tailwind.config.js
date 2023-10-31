/** @type {import('tailwindcss').Config} */
module.exports = {
  darkMode: 'media',
  content: ["./internal/templates/**/*.{html,js}"],
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
            '--tw-prose-quote-borders': '#3a3737',
            'blockquote p:first-of-type::before': false,
            'blockquote p:first-of-type::after': false,
            color: '#3a3737',
            blockquote: {
              color: '#3a3737',
            },
            h2: {
              color: '#3a3737',
            },
            h3: {
              color: '#3a3737',
            },
            pre: {
              backgroundColor: 'black',
              color: 'silver',
            },
            code: {
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
            '--tw-prose-quote-borders': '#fbf1c7',
            color: '#fbf1c7',
            blockquote: {
              color: '#fbf1c7',
            },
            h2: {
              color: '#fbf1c7',
            },
            h3: {
              color: '#fbf1c7',
            },
            pre: {
              backgroundColor: '#32302f',
              color: '#fbf1c7',
            },
            code: {
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
