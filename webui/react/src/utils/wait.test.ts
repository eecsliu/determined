import { RecordKey } from 'shared/types';
import { generateUUID } from 'shared/utils/string';
import { CommandState, CommandTask, CommandType } from 'types';

import * as utils from './wait';

const UUID = {
  [CommandType.JupyterLab]: generateUUID(),
  [CommandType.Shell]: generateUUID(),
  [CommandType.TensorBoard]: generateUUID(),
};
const UUID_REGEX = '[a-z\\d-]+';
const SHARED_TASK: Partial<CommandTask> = {
  misc: undefined,
  url: undefined,
};
const COMMAND_TASK: Record<RecordKey, CommandTask> = {
  [CommandType.JupyterLab]: {
    ...SHARED_TASK,
    displayName: 'lee sung hoon',
    id: UUID[CommandType.JupyterLab],
    name: 'JupyterLab (directly-sharp-penguin)',
    resourcePool: 'compute-pool',
    serviceAddress: `/proxy/${UUID[CommandType.JupyterLab]}`,
    startTime: '2022-08-01T00:12:24Z',
    state: CommandState.Pending,
    type: CommandType.JupyterLab,
    userId: 34,
  },
  [CommandType.Shell]: {
    ...SHARED_TASK,
    displayName: 'lee sung hoon',
    id: UUID[CommandType.Shell],
    name: 'Shell (jolly-well-pheasant)',
    resourcePool: 'compute-pool',
    serviceAddress: `/proxy/${UUID[CommandType.Shell]}`,
    startTime: '2022-07-29T24:00:12Z',
    state: CommandState.Terminated,
    type: CommandType.Shell,
    userId: 34,
  },
  [CommandType.TensorBoard]: {
    ...SHARED_TASK,
    displayName: 'racer kim',
    id: UUID[CommandType.TensorBoard],
    name: 'TensorBoard (ambiguously-happy-ear)',
    resourcePool: 'aux-pool',
    serviceAddress: `/proxy/${UUID[CommandType.TensorBoard]}`,
    startTime: '2022-08-02T12:24:00Z',
    state: CommandState.Running,
    type: CommandType.TensorBoard,
    userId: 16,
  },
};

describe('Wait Page Utilities', () => {
  describe('commandToEventUrl', () => {
    const REGEX: Record<RecordKey, RegExp> = {
      [CommandType.JupyterLab]: new RegExp(`notebooks/${UUID_REGEX}/events`, 'i'),
      [CommandType.TensorBoard]: new RegExp(`tensorboard/${UUID_REGEX}/events`, 'i'),
      [CommandType.Shell]: new RegExp('^$', 'i'),
    };

    it('should convert task to event url', () => {
      for (const [ type, task ] of Object.entries(COMMAND_TASK)) {
        expect(utils.commandToEventUrl(task)).toMatch(REGEX[type]);
      }
    });
  });

  describe('openCommand', () => {
    let globalOpen: typeof global.open;
    const windowOpen = jest.fn();

    beforeAll(() => {
      // Preserve the original `global.open`.
      globalOpen = global.open;
      global.open = windowOpen;
    });

    afterAll(() => {
      // Restore `global.open` to original function.
      global.open = globalOpen;
    });

    it('should open window for JupyterLab task', () => {
      expect(windowOpen).not.toHaveBeenCalled();
      utils.openCommand(COMMAND_TASK[CommandType.JupyterLab]);
      expect(windowOpen).toHaveBeenCalled();
    });

    it('should open window for TensorBoard task', () => {
      expect(windowOpen).not.toHaveBeenCalled();
      utils.openCommand(COMMAND_TASK[CommandType.TensorBoard]);
      // TODO: Expand this to use `toHaveBeenCalledWith`.
      expect(windowOpen).toHaveBeenCalled();
    });

    it('should throw error for tasks that are not JupyterLabs or TensorBoards', () => {
      expect(() => {
        utils.openCommand(COMMAND_TASK[CommandType.Shell]);
      }).toThrow(utils.CANNOT_OPEN_COMMAND_ERROR);
    });
  });

  describe('waitPageUrl', () => {
    const REGEX: Record<RecordKey, RegExp> = {
      [CommandType.JupyterLab]: new RegExp(`wait/${CommandType.JupyterLab}/${UUID_REGEX}`, 'i'),
      [CommandType.TensorBoard]: new RegExp(`wait/${CommandType.TensorBoard}/${UUID_REGEX}`, 'i'),
    };

    it('should convert task to wait page url', () => {
      expect(utils.waitPageUrl(COMMAND_TASK[CommandType.JupyterLab]))
        .toMatch(REGEX[CommandType.JupyterLab]);

      expect(utils.waitPageUrl(COMMAND_TASK[CommandType.TensorBoard]))
        .toMatch(REGEX[CommandType.TensorBoard]);
    });

    it('should throw error for tasks that are not JupyterLabs or TensorBoards', () => {
      expect(() => {
        utils.waitPageUrl(COMMAND_TASK[CommandType.Shell]);
      }).toThrow(utils.CANNOT_OPEN_COMMAND_ERROR);
    });
  });
});