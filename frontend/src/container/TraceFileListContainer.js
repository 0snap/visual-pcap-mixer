import { connect } from 'react-redux'
import TraceFileList from '../representation/TraceFileList';

function getVisibleTraceFiles(traceFiles) {
    return traceFiles;
}

const mapStateToProps = state => {
    return {
        traceFiles: getVisibleTraceFiles(state.traceFiles)
    }
}

const mapDispatchToProps = {}

const TraceFileListContainer = connect(
    mapStateToProps,
    mapDispatchToProps
)(TraceFileList)

export default TraceFileListContainer