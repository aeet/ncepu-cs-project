import { Environment } from '@yelon/theme';

export const environment = {
  production: true,
  useHash: true,
  path: 'http://localhost:8080',
  api: {
    baseUrl: './',
    refreshTokenEnabled: true,
    refreshTokenType: 'auth-refresh'
  }
} as Environment;
