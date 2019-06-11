import { connect } from 'react-redux'
import MsaTimeline from '../representation/MsaTimeline';
import {changeTimelineOrder, removeMsaEntry, removeMsaDay, addMsaDay, addMsaEntry} from '../actions';

const mapStateToProps = state => {
    return {
        msaTimeline: state.msaTimeline,
        attacks: state.attacks,
        traceFiles: state.traceFiles,
    }
}

const mapDispatchToProps = dispatch => {
    return {
        dragAndDrop: (draggedDay, droppedOnDay) => {
            dispatch(changeTimelineOrder(draggedDay, droppedOnDay));
        },
        entryDoubleClick: (dayId, entryId) => {
            dispatch(removeMsaEntry(dayId, entryId));
        },
        dayDoubleClick: (id) => {
            dispatch(removeMsaDay(id));
        },
        createDay: () => {
            dispatch(addMsaDay());
        },
        addMsaEntry: (dayId, entryId, entryType) => {
            dispatch(addMsaEntry(dayId, {id: entryId, type: entryType} ));
        }
    };
};

const MsaTimelineContainer = connect(
    mapStateToProps,
    mapDispatchToProps
)(MsaTimeline)

export default MsaTimelineContainer