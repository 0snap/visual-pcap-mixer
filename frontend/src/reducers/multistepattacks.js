const defaultState = {
    waiting: false,
    success: false,
    error: false,
    knownAttacks: {},
};

function multistepattacks(state = defaultState, action) {
    switch (action.type) {
    case 'CREATE_MSA':
        return {...state, waiting: true, success: false, error: false};
    case 'CREATE_MSA_SUCCESS':
        const knownAttacks = {...state.knownAttacks};
        knownAttacks[action.msa.name] = action.msa;
        return {knownAttacks, waiting: false, success: true, error: false};
    case 'CREATE_MSA_ERROR':
        return {...defaultState, error: true};
    case 'LOAD_MSA_SUCCESS':
        return {knownAttacks: action.multistepattacks, waiting: false, success: false, error: false};
    default:
        return state
    }
}
export default multistepattacks