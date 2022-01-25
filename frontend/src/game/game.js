import Phaser from 'phaser';
import BootScene from './scenes/BootScene';

function launch(containerId) {
  return new Phaser.Game({
    type: Phaser.AUTO,
    backgroundColor: '#96bb58',
    scale: {
      width: 1920,
      height: 1080,
      parent: containerId,
      mode: Phaser.Scale.FIT,
      autoCenter: Phaser.Scale.CENTER_BOTH,
    },
    physics: {
      default: 'arcade',
      arcade: {
        gravity: { y: 0 },
        debug: false,
      },
    },
    scene: [BootScene],
  });
}

export default launch;
export { launch };
