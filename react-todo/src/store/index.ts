import { create } from 'zustand';

type EditedTask = {
    id: number;
    title: string;
}

type State = {
    editedTask: EditedTask
    updateEditedTask: (payload: EditedTask) => void
}

const userStore = create<State>((set) => ({
    editedTask: { id: 0, title: '' },
    updateEditedTask: (payload) =>
      set({
        editedTask: payload,
      }),  
}));

export default userStore;
