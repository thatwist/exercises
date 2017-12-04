using System.Collections;
using System.Collections.Generic;
using UnityEngine;

public class TurnToParticles : MonoBehaviour {

	public GameObject particlePrefab;

	void OnCollisionEnter(Collision collision) {
		Instantiate (particlePrefab, this.gameObject.transform.position, Quaternion.identity);
		Destroy (this.gameObject);
	}
}
