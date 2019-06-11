import MsaApi from '../api/msaApi';
import { loadAttacksSuccess, loadTraceFilesSuccess, loadMultiStepAttacksSuccess, createMsaSuccess, createMsaError } from '.';

export function loadAll() {
    return function(dispatch) {
        return MsaApi.getAll()
            .then(res => {
                dispatch(loadAttacksSuccess(res.attacks));
                dispatch(loadMultiStepAttacksSuccess(res.multistepattacks));
                return dispatch(loadTraceFilesSuccess(res.traceFiles))
            });
    };
}

export function createMsa(msa) {
    return function(dispatch) {
        return MsaApi.createMsa(msa)
            .then(res => dispatch(createMsaSuccess(res)))
            .catch(err => {
                console.error('API error creating MSA', err);
                dispatch(createMsaError(err.message));
            })
    }
}