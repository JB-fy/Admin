/* eslint-env node */
require('@rushstack/eslint-patch/modern-module-resolution')

module.exports = {
  root: true,
  'extends': [
    'plugin:vue/vue3-essential',
    'eslint:recommended',
    '@vue/eslint-config-typescript',
    '@vue/eslint-config-prettier/skip-formatting'
  ],
  parserOptions: {
    ecmaVersion: 'latest'
  },
  /* globals: {
    ElMessage: 'readonly',
  }, */
  ignorePatterns: ['public/**'],
  /* rules: {
    'vue/multi-word-component-names': 'off'
  }, */
  overrides: [
    {
      files: ['src/views/**/*.vue', 'src/layout/**/*.vue'],
      rules: {
        'vue/multi-word-component-names': 'off'
      }
    }
  ]
}
