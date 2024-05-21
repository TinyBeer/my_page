import React, { useState, useEffect } from 'react';
import { useDispatch, useSelector } from 'react-redux';
import { Col, Row, Skeleton, Image } from 'antd';
import { refreshToken } from './store/module/tokenStore';
import { useNavigate } from 'react-router';
import { fetchVideoList } from './store/module/cinemaStore';

export default function App({ displayAlert }) {
  const [loginStatus, setLoginStatus] = useState(false);
  const [loading, setloading] = useState(true);
  const navigate = useNavigate();
  const dispatch = useDispatch();
  useEffect(() => {
    if (!loginStatus) {
      refreshToken(
        () => setLoginStatus(true),
        () => navigate('/login')
      );
    }

    const run = async () => {
      dispatch(
        fetchVideoList(
          () => setloading(false),
          () => navigate('/login')
        )
      );
    };
    run();
  }, [dispatch]);

  const videoList = useSelector((state) => state.cinema.videoList);
  return (
    <Row
      gutter={{
        xs: 8,
        sm: 16,
        md: 24,
        lg: 32,
      }}
    >
      {videoList &&
        videoList.length > 0 &&
        videoList.map((item, index) => {
          const key = `col-${index}`;
          return (
            <Col
              key={key}
              xs={{
                flex: '40%',
              }}
              sm={{
                flex: '35%',
              }}
              md={{
                flex: '32%',
              }}
              lg={{
                flex: '30%',
              }}
              xl={{
                flex: '25%',
              }}
            >
              <div style={{ padding: '5px' }}>
                {loading ? (
                  <div>
                    <Skeleton.Image loading active />
                    <Skeleton
                      active={true}
                      paragraph={{
                        rows: 1,
                      }}
                    ></Skeleton>
                  </div>
                ) : (
                  <Image
                    width={180}
                    height={240}
                    preview={{
                      destroyOnClose: true,
                      imageRender: () => (
                        <video
                          // muted
                          width='80%'
                          controls
                          src={
                            'http://127.0.0.1:9999/static/video/movie/' +
                            item.name
                          }
                        />
                      ),
                      toolbarRender: () => null,
                    }}
                    src={`http://127.0.0.1:9999/static/img/` + item.image}
                  />
                )}
              </div>
            </Col>
          );
        })}
    </Row>
  );
}