import Phaser from 'phaser';
import BootScene from './scenes/BootScene';
import MenuScene from './scenes/MenuScene';
import PlayScene from './scenes/PlayScene';

function launch(containerId) {
  return new Phaser.Game({
    type: Phaser.CANVAS,
    backgroundColor: '#96bb58',
    scale: {
      width: 1920,
      height: 1080,
      parent: containerId,
      mode: Phaser.Scale.FIT,
      autoCenter: Phaser.Scale.CENTER_BOTH,
    },
    fps: {
      target: 60,
      min: 60,
      forceSetTimeOut: true,
    },
    physics: {
      default: 'arcade',
      arcade: {
        gravity: { y: 0 },
        debug: false,
      },
    },
    scene: [BootScene, PlayScene, MenuScene],
  });
}

export default launch;
export { launch };
