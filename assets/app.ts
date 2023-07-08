import * as PIXI from 'pixijs'

document.addEventListener('DOMContentLoaded', () => {
  const app = new PIXI.Application <HTMLCanvasElement>({ background: '#000', resizeTo: window });
  document.body.appendChild(app.view);
})
