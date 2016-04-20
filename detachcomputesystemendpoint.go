package hcsshim

import "github.com/Sirupsen/logrus"

// DetachComputeSystemEndpoint joins an endpoint to a container.
func DetachComputeSystemEndpoint(id string, epid string) error {

	title := "HCSShim::DetachComputeSystemEndpoint"
	logrus.Debugln(title+" id=%s, epid=%s", id, epid)

	err := detachComputeSystemEndpoint(id, epid)
	if err != nil {
		err = makeErrorf(err, title, "id=%s epid=%s", id, epid)
		logrus.Error(err)
		return err
	}

	logrus.Debugf(title+"- succeeded %s", id)
	return nil
}
