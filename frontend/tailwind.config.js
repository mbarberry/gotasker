module.exports = {
  purge: ['./index.html', './src/**/*.{vue,js,ts,jsx,tsx}'],
  darkMode: false,
  theme: {
    extend: {
      colors: {
        gopher: '#00ADD8',
      },
    },
  },
  variants: {
    cursor: ({ after }) => after(['disabled']),
  },
  plugins: [],
};
