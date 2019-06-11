function msaTimeline(state = [], action) {
    const newState = state.slice();
    switch (action.type) {
    case 'ADD_MSA_DAY':
        newState.push([]);
        return newState;
    case 'REMOVE_MSA_DAY':
        newState.splice(action.id, 1);
        return newState;
    case 'ADD_MSA_ENTRY':
        newState[action.dayId].push(action.entry);
        return newState;
    case 'REMOVE_MSA_ENTRY':
        newState[action.dayId] = newState[action.dayId].filter(entry => entry.id !== action.entryId);
        return newState;
    case 'LOAD_MSA_TO_TIMELINE':
        return action.msa.timeline;
    case 'CREATE_MSA_SUCCESS':
        return action.msa.timeline;
    case 'CHANGE_TIMELINE_ORDER':
        const elemA = newState.splice(action.a, 1)[0];
        newState.splice(action.b, 0, elemA);
        return newState;
    default:
        return state
    }
}
export default msaTimeline;