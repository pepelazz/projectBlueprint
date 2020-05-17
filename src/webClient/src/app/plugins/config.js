export default {
  appName: 'CompanyName',
  uiAppName: 'CompanyName',
  apiUrl: () => process.env.NODE_ENV === 'development' ? 'http://localhost:3091' : 'https://example.ru',
  wsUrl: () => process.env.NODE_ENV === 'development' ? 'ws://localhost:3091' : 'wss://example.ru',
  isEmailAuth: {
    firstName: true,
    lastName: true,
  },
  logoSrc: 'https://cdn.pixabay.com/photo/2017/05/05/00/15/kokopelli-2285538_960_720.png',
  dadataToken: '1cf3a086e3dbe1306ed142d2b5fbc1b324b8eb60',
  // yandexMetrikaId: 54433825,
  breadcrumbIcons: {
    legal_entity: 'far fa-file-alt'
  },
  tablesForTask: {},
}
