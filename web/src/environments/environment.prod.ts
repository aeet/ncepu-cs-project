import { Environment } from '@yelon/theme';

export const environment = {
  production: true,
  useHash: true,
  path: '',
  api: {
    baseUrl: './',
    refreshTokenEnabled: true,
    refreshTokenType: 'auth-refresh'
  }
} as Environment;
