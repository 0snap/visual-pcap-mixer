function replacements(state = [], action) {
    const newState = state.slice();
    switch (action.type) {
    case 'ADD_REPLACEMENT':
        newState.push( { ipA: action.ipA, ipB: action.ipB });
        return newState;
    case 'REMOVE_REPLACEMENT':
        newState.splice(action.id);
        return newState;
    default:
        return state
    }
}
export default replacements