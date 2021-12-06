import React, { useEffect, useState } from 'react';
import DescriptionVideoCard from '../descriptionVideoCard/DescriptionVideoCard';
import SpotlightVideoCard from '../spotlightVideoCard/SpotlightVideoCard';
import PlusVideosCard from '../plusVideosCard/PlusVideosCard';
import * as YoutubeService from '../../services/youtubeServices/YoutubeServices';

function MainPage() {
    const [state, setState] = useState({
        spotlightVideoUrl: null,
        videoDescription: null,
        relatedVideos: [],
        videoTitle: null
    });

    useEffect(() => {
        YoutubeService.getUnsubscribedTrailer().then((response) => {
            setState({
                spotlightVideoUrl: response.data.items[0].brandingSettings.channel.unsubscribedTrailer,
            },
            () => YoutubeService.getRelatedVideos(state.spotlightVideoUrl).then((response) => {
                setState({
                    relatedVideos: response.data.items,
                })
                YoutubeService.getFullDescriptionVideo(state.spotlightVideoUrl).then((response) => {
                    setState({
                        videoDescription: response.data.items[0].snippet.description,
                        videoTitle: response.data.items[0].snippet.title,
                    })
                })
            }))
        })
    })

    function changeVideo(newVideo) {
        setState({
            spotlightVideoUrl: newVideo,
        })

        YoutubeService.getRelatedVideos(newVideo).then((response) => {
            setState({
                relatedVideos: response.data.items,
            })
            YoutubeService.getFullDescriptionVideo(newVideo).then((response) => {
                setState({
                    videoDescription: response.data.items[0].snippet.description,
                })
            })
        })
    }

    return (
        <div style={styles.mainDiv}>
            <div style={styles.leftCard}>
                <SpotlightVideoCard
                spotlightUrl={state.spotlightVideoUrl}
                />
                <div style={styles.outDescriptionDiv}>
                    <div style={styles.descriptionDiv}>
                        <DescriptionVideoCard
                        videoTitle={state.videoTitle}
                        videoDescription={state.videoDescription}
                        />
                    </div>
                </div>
            </div>
            <div style={styles.rightCard}>
                <PlusVideosCard
                changeVideo={changeVideo}
                arrayVideosRelated={state.relatedVideos}
                />
            </div>
        </div>
    );
}

const styles = {
    mainDiv: {
        display: "flex",
        flexWrap: "wrap",
        justifyContent: "center",
    },
    descriptionDiv: {
        maxWidth: 560,
    },
    outDescriptionDiv: {
        display: "flex",
        justifyContent: "center"
    },
    leftCard: {
        width: "50%",
        minWidth: 345,
    },
    rightCard: {
        width: "20%",
        minWidth: 345,
    },
}


export default MainPage;