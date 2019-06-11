function parseAttackDictionary(attacks, defaultAdded = false) {
    return Object.entries(attacks).reduce((acc, entry) => {
        const [id, atk] = entry;
        atk.added = defaultAdded;
        acc[id] = atk;
        return acc;
    }, {});
}

function attacks(state = {}, action) {
    const newState = {...state};
    switch (action.type) {
    case 'ADD_MSA_ENTRY':
        if (action.entry.type === 'attack') newState[action.entry.id].added = true;
        return newState;
    case 'REMOVE_MSA_ENTRY':
        if (newState[action.entryId]) newState[action.entryId].added = false;
        return newState;
    case 'LOAD_ATTACKS_SUCCESS':
        return parseAttackDictionary(action.attacks);
    case 'LOAD_MSA_SUCCESS':
        const loadedAttacksFromMsa = Object.entries(action.multistepattacks).reduce((acc, entry) => {
            const msa = entry[1];
            return {...acc, ...parseAttackDictionary(msa.attacks)};
        }, {});
        return {...newState, ...loadedAttacksFromMsa};
    case 'CREATE_MSA_SUCCESS':
        return {...newState, ...parseAttackDictionary(action.msa.attacks, true)};
    case 'LOAD_MSA_TO_TIMELINE':
        return parseAttackDictionary(state);
    default:
        return state
    }
}
export default attacks