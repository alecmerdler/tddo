export { Mode, Point, State, start };

enum Mode {
  Up = 'Up',
  Down = 'Down',
}
type Point = {x: number, y: number};

type State = {Mode: Mode, Point: Point};

const start: State = {Mode: Mode.Up, Point: {x: 0, y: 0}};
