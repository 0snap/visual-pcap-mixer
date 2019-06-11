import { combineReducers } from 'redux'
import attacks from './attacks'
import traceFiles from './traceFiles'
import replacements from './replacements'
import msaTimeline from './msaTimeline'
import multistepattacks from './multistepattacks'

export default combineReducers({
    attacks,
    traceFiles,
    replacements,
    multistepattacks,
    msaTimeline,
})