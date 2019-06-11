// attacks ---------------------------------------------------------------------

export function loadAttacksSuccess(attacks) {
    return {type: 'LOAD_ATTACKS_SUCCESS', attacks};
}

// trace files -----------------------------------------------------------------

export function loadTraceFilesSuccess(traceFiles) {
    return {type: 'LOAD_TRACEFILES_SUCCESS', traceFiles};
}

// replacements ----------------------------------------------------------------
export function addReplacement(ipA, ipB) {
    return {type: 'ADD_REPLACEMENT', ipA, ipB};
}

export function removeReplacement(id) {
    return {type: 'REMOVE_REPLACEMENT', id};
}

// msa timeline ----------------------------------------------------------------

export function addMsaEntry (dayId, entry) {
    return {type: 'ADD_MSA_ENTRY', dayId, entry};
}

export function removeMsaEntry (dayId, entryId) {
    return {type: 'REMOVE_MSA_ENTRY', dayId, entryId};
}

export function changeTimelineOrder(a, b) {
    return {type: 'CHANGE_TIMELINE_ORDER', a, b}
}

export function addMsaDay () {
    return {type: 'ADD_MSA_DAY'};
}

export function removeMsaDay (id) {
    return {type: 'REMOVE_MSA_DAY', id};
}

// msa API actions -------------------------------------------------------------

export function createMsa() {
    return {type: 'CREATE_MSA'};
}

export function createMsaSuccess(msa) {
    return {type: 'CREATE_MSA_SUCCESS', msa};
}

export function createMsaError(msg) {
    return {type: 'CREATE_MSA_ERROR', msg};
}

// multistepattack action (i.e. interacting with already created MSAs) ---------

export function loadMultiStepAttacksSuccess(multistepattacks) {
    return {type: 'LOAD_MSA_SUCCESS', multistepattacks};
}
export function loadMsaToTimeline(msa) {
    return {type: 'LOAD_MSA_TO_TIMELINE', msa};
}