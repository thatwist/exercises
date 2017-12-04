using System.Collections;
using System.Collections.Generic;
using UnityEngine;

public class PlayerController : MonoBehaviour {

	public int speed; 
	public GameObject ball;

	void Update () {
		float horizontalAxis = Input.GetAxis ("Horizontal");
		transform.position = 
			new Vector3(
				Mathf.Clamp(transform.position.x + horizontalAxis * speed, -8f, 8f), 
				transform.position.y, 
				transform.position.z
			);

		if (Input.GetKeyDown ("space")) {
			ball.transform.parent = null;
			Rigidbody rb = ball.GetComponent<Rigidbody> ();
			rb.isKinematic = false;
			rb.AddForce (new Vector3 (800f, 800f, 0.0f));
		}
	}
}
