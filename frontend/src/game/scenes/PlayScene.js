import { Scene, Actions, Math as PhaserMath } from 'phaser';
import axios from 'axios';
import background from '@/game/assets/background.svg';
import snakeHead from '@/game/assets/head.svg';
import snakeBody from '@/game/assets/body.svg';
import greenApple from '@/game/assets/green-apple.svg';
import redApple from '@/game/assets/red-apple.svg';
import horizontalObstacle from '@/game/assets/horizontal-obstacle.svg';
import verticalObstacle from '@/game/assets/vertical-obstacle.svg';
import eatGreenSound from '@/game/assets/eat-green.ogg';
import eatRedSound from '@/game/assets/eat-red.ogg';
import deadSound from '@/game/assets/dead-sound.ogg';

const WIDTH = 1920;
const HEIGHT = 1080;

const BACKGROUND = 'background';
const SNAKE_HEAD = 'snake-head';
const SNAKE_BODY = 'snake-body';
const GREEN_APPLE = 'green-apple';
const RED_APPLE = 'red-apple';

const HORIZONTAL_OBSTACLE = 'horizontal-obstacle';
const VERTICAL_OBSTACLE = 'vertical-obstacle';

const EAT_GREEN_SOUND = 'eat-green-sound';
const EAT_RED_SOUND = 'eat-red-sound';

const DEAD_SOUND = 'dead-sound';

export default class PlayScene extends Scene {
  constructor() {
    super({ key: 'PlayScene' });
  }

  init(data) {
    this.backgroundMusic = data.music;
    this.score = 0;
    this.eaten = 0;
    this.alive = true;
  }

  preload() {
    this.load.image(BACKGROUND, background);
    this.load.image(SNAKE_HEAD, snakeHead);
    this.load.image(SNAKE_BODY, snakeBody);
    this.load.image(GREEN_APPLE, greenApple);
    this.load.image(RED_APPLE, redApple);
    this.load.image(HORIZONTAL_OBSTACLE, horizontalObstacle);
    this.load.image(VERTICAL_OBSTACLE, verticalObstacle);

    this.load.audio(EAT_GREEN_SOUND, eatGreenSound);
    this.load.audio(EAT_RED_SOUND, eatRedSound);
    this.load.audio(DEAD_SOUND, deadSound);
  }

  create() {
    this.cameras.main.fadeIn(1000);
    this.tweens.add({
      targets: this.backgroundMusic,
      volume: 0,
      duration: 20000,
    });

    this.background = this.add.tileSprite(0, 0, 1920, 1920, BACKGROUND).setOrigin(0);
    const boundWidth = this.background.width;
    const boundHeight = this.background.height;
    this.physics.world.setBounds(0, 0, boundWidth, boundHeight, true, true, true, true);

    this.body = this.add.group({ key: SNAKE_BODY, frameQuantity: 2 });

    this.tail = this.body.getLast(true);

    this.head = this.physics.add.image(WIDTH / 2, HEIGHT / 2, SNAKE_HEAD);
    this.head.setCollideWorldBounds(true);
    this.head.setDepth(1);

    this.input.on('pointermove', (pointer) => {
      if (this.alive === true) {
        this.physics.moveTo(this.head, pointer.worldX, pointer.worldY, 500);
      }
    }, this);

    this.cameras.main.startFollow(this.head, this.cameras.main.FOLLOW_LOCKON, 0.05, 0.05);

    this.obstacles = this.physics.add.group();

    this.food = this.physics.add.group();

    this.time.addEvent({
      delay: 2000,
      callback: this.generateFood,
      args: [GREEN_APPLE],
      callbackScope: this,
      loop: true,
    });
    this.time.addEvent({
      delay: 5000,
      callback: this.generateFood,
      args: [RED_APPLE],
      callbackScope: this,
      loop: true,
    });

    this.physics.add.overlap(this.head, this.food, this.eat, null, this);
    this.physics.add.collider(this.head, this.obstacles, this.die, null, this);
  }

  generateFood(type) {
    const add = type === GREEN_APPLE || Boolean(Math.random() <= 0.5);

    if (this.food.countActive() > 19 || !add) return;

    const randomX = PhaserMath.Between(this.background.width * 0.1, this.background.width * 0.9);
    const randomY = PhaserMath.Between(this.background.height * 0.1, this.background.height * 0.9);

    this.food.create(randomX, randomY, type);
  }

  eat(head, food) {
    this.eaten += 1;

    if (food.texture.key === GREEN_APPLE) {
      this.body.create(-10000, 0, SNAKE_BODY);
      this.sound.play(EAT_GREEN_SOUND);
      this.score += 1;
    } else {
      for (let i = 0; i < this.bodyparts.length / 2; i += 1) {
        this.tail = this.body.getLast(true);
        if (this.tail != null) {
          this.tail.destroy(true);
          this.sound.play(EAT_RED_SOUND);
          this.score -= 1;
        }
      }
    }

    if (this.bodyparts.length === 1) {
      this.die();
    }

    if (this.eaten % 10 === 0 && this.eaten !== 0) {
      this.generateObstacle();
    }

    food.destroy(true);
  }

  generateObstacle() {
    const randomX = PhaserMath.Between(this.background.width * 0.1, this.background.width * 0.9);
    const randomY = PhaserMath.Between(this.background.height * 0.1, this.background.height * 0.9);

    const angle = Math.random() <= 0.5 ? VERTICAL_OBSTACLE : HORIZONTAL_OBSTACLE;

    const obstacle = this.obstacles.create(randomX, randomY, angle);
    obstacle.setBounce(1);
    obstacle.setCollideWorldBounds(true);

    if (angle === VERTICAL_OBSTACLE) {
      obstacle.setVelocityY(450);
    } else {
      obstacle.setVelocityX(450);
    }
  }

  die(head, obstacle) {
    this.alive = false;
    this.sound.play(DEAD_SOUND, { volume: 0.3 });

    if (obstacle !== undefined) {
      obstacle.setBounce(0);
      obstacle.setVelocity(0, 0);
    }

    this.head.setVelocity(0, 0);

    this.cameras.main.fadeOut(500);

    if (this.score > 0) {
      this.updateScore();
    }

    setTimeout(() => {
      this.scene.start('MenuScene', { score: this.score, music: this.backgroundMusic });
    }, 500);
  }

  updateScore() {
    const config = {
      method: 'put',
      url: `${process.env.VUE_APP_BASE_URL}score`,
      headers: {
        Authorization: `Bearer ${localStorage.getItem('token')}`,
        'Content-Type': 'application/json',
      },
      data: {
        max_score: this.score,
      },
    };

    axios(config)
      .catch((err) => {
        console.log(err);
      });
  }

  update() {
    this.bodyparts = this.body.getChildren();
    if (this.bodyparts.length > 0) {
      Actions.ShiftPosition(this.bodyparts, this.head.x, this.head.y);
    }
  }
}
