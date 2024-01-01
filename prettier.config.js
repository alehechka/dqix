/** @type {import("prettier").Config} */
const config = {
  trailingComma: 'es5',
  tabWidth: 2,
  semi: true,
  singleQuote: true,
  length: 80,
  plugins: ['prettier-plugin-tailwindcss'],
  overrides: [
    {
      files: ['*.templ'],
      options: {
        parser: 'tailwindcss',
      },
    },
  ],
};

module.exports = config;
