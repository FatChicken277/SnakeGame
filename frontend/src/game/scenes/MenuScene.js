import { Scene } from 'phaser';
import menuBackground from '@/game/assets/menu-background.svg';

const HEIGHT = 1080;
const WIDTH = 1920;

const MENU_BACKGROUND = 'menu-background';

export default class MenuScene extends Scene {
  constructor() {
    super({ key: 'MenuScene' });
  }

  init(data) {
    this.score = data.score < 0 ? 0 : data.score;
    this.backgroundMusic = data.music;
  }

  preload() {
    this.load.image(MENU_BACKGROUND, menuBackground);
  }

  create() {
    this.cameras.main.fadeIn(1000);
    this.tweens.add({
      targets: this.backgroundMusic,
      volume: 0.5,
      duration: 10000,
    });

    this.background = this.add.image(0, 0, MENU_BACKGROUND).setScale(3).setOrigin(0);

    const scoreText = this.add.text(WIDTH / 2, HEIGHT / 3, `SCORE: ${this.score}`, { font: '100px VT323', shadow: '10px' }).setOrigin(0.5);
    scoreText.setShadow(5, 5, 'rgba(0,0,0,0.5)', 15);

    const playAgainText = this.add.text(WIDTH / 2, HEIGHT / 1.8, 'PLAY AGAIN', { font: '50px VT323' }).setOrigin(0.5);
    playAgainText.setShadow(5, 5, 'rgba(0,0,0,0.5)', 15);
    playAgainText.setInteractive();
    playAgainText.on('pointerdown', () => {
      this.cameras.main.fadeOut(500);
      setTimeout(() => {
        this.scene.start('PlayScene');
      }, 500);
    });

    const menuText = this.add.text(WIDTH / 2, HEIGHT / 1.5, 'MENU', { font: '30px VT323' }).setOrigin(0.5);
    menuText.setShadow(5, 5, 'rgba(0,0,0,0.5)', 15);
    menuText.setInteractive();
    menuText.on('pointerdown', () => {
      this.cameras.main.fadeOut(500);
      setTimeout(() => {
        this.scene.start('BootScene', { music: this.backgroundMusic });
      }, 500);
    });
  }
}
