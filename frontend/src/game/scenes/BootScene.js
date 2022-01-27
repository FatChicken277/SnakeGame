import { Scene } from 'phaser';
import bootBackground from '@/game/assets/boot-background.svg';
import backgroundMusic from '@/game/assets/background-music.ogg';
import fullscreen from '@/game/assets/fullscreen.png';

const WIDTH = 1920;
const HEIGHT = 1080;

const BOOT_BACKGROUND = 'boot-background';
const BACKGROUND_MUSIC = 'background-music';
const FULLSCREEN = 'fullscreen';

export default class BootScene extends Scene {
  constructor() {
    super({ key: 'BootScene' });
  }

  init(data) {
    this.backgroundMusic = data.music;
  }

  preload() {
    this.load.image(BOOT_BACKGROUND, bootBackground);
    this.load.audio(BACKGROUND_MUSIC, backgroundMusic);
    this.load.spritesheet(FULLSCREEN, fullscreen, { frameWidth: 512, frameHeight: 512 });
  }

  create() {
    this.cameras.main.fadeIn(500);

    if (this.backgroundMusic === undefined) {
      this.backgroundMusic = this.sound.add(BACKGROUND_MUSIC, { volume: 0, loop: true });
      this.backgroundMusic.play();
      this.tweens.add({
        targets: this.backgroundMusic,
        volume: 0.5,
        duration: 10000,
      });
    }

    this.background = this.add.image(0, 0, BOOT_BACKGROUND).setScale(3).setOrigin(0);

    const fsButton = this.add.image(1920 - 16, 16, FULLSCREEN, 0).setOrigin(1, 0).setInteractive();
    fsButton.setScale(0.1);

    if (this.scale.isFullscreen) {
      fsButton.setFrame(1);
    } else {
      fsButton.setFrame(0);
    }

    fsButton.on('pointerup', () => {
      if (this.scale.isFullscreen) {
        fsButton.setFrame(0);
        this.scale.stopFullscreen();
      } else {
        fsButton.setFrame(1);
        this.scale.startFullscreen();
      }
    }, this);

    const logoText = this.add.text(WIDTH / 2, HEIGHT / 3, 'SNAKEE', { font: '150px VT323', shadow: '10px' }).setOrigin(0.5);
    logoText.setShadow(5, 5, 'rgba(0,0,0,0.5)', 15);

    const playText = this.add.text(WIDTH / 2, HEIGHT / 1.8, 'PLAY', { font: '50px VT323' }).setOrigin(0.5);
    playText.setShadow(5, 5, 'rgba(0,0,0,0.5)', 15);
    playText.setInteractive();
    playText.on('pointerdown', () => {
      this.cameras.main.fadeOut(500);
      setTimeout(() => {
        this.scene.start('PlayScene', { music: this.backgroundMusic });
      }, 500);
    });
  }
}
