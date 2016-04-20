package hcsshim

import "github.com/Sirupsen/logrus"

// AttachComputeSystemEndpoint joins an endpoint to a container.
func AttachComputeSystemEndpoint(id string, epid string) error {

	title := "HCSShim::AttachComputeSystemEndpoint"
	logrus.Debugln(title+" id=%s, epid=%s", id, epid)

	err := attachComputeSystemEndpoint(id, epid)
	if err != nil {
		err = makeErrorf(err, title, "id=%s epid=%s", id, epid)
		logrus.Error(err)
		return err
	}

	logrus.Debugf(title+"- succeeded %s", id)
	return nil
}
